package snowflake

import (
	"github.com/tursom/GoCollections/util/time"

	"github.com/tursom/GoCollections/lang"
	"github.com/tursom/GoCollections/lang/atomic"
	"github.com/tursom/GoCollections/util/b62"
)

/**
 * 被改造过的雪花ID
 * |---------|------------|-----------|------------|
 * | 空位(1) | 时间戳(43) | 自增位(7) | 机器ID(13) |
 * |---------|------------|-----------|------------|
 * 时间戳仅在初始化时使用, 与序列号接壤, 这样做可以避免某一时间内大量请求导致的ID爆炸
 * 当前方案在ID溢出时, 溢出的数据会让时间戳 +1.
 * 这样做, 只要节点不重启或者在重启前平均QPS没有超标, 重启后分配的ID仍能唯一
 *
 * 当前被调试为平均每毫秒可以相应 128 个消息
 * 如果平均每个人 10 秒发一条消息, 1 秒 128 条消息大约要 1280 人, 1 毫秒 128 条消息就大约需要 128W 用户了
 * 单点无法应付如此巨量的并发, ID生成器保证性能过剩
 *
 * 当前最多支持 8192 个节点同时上线, 未来如果节点数超过了 8192 个, 也可以以ID生成的最晚时间为代价提升节点最高数量
 */

const (
	incrementBase uint64 = 0x2000
)

type (
	Snowflake interface {
		New() uint64
		NewStr() string
		Close()
	}

	snowflake struct {
		nodeId          uint16
		incrementLength int
		seed            atomic.UInt64
		closer          lang.SendChannel[struct{}]
	}
)

func New(nodeId uint16) Snowflake {
	if uint64(nodeId) >= incrementBase {
		panic(nodeId)
	}

	closer := make(chan struct{})

	s := &snowflake{
		nodeId:          nodeId,
		incrementLength: 7,
		closer:          lang.RawChannel[struct{}](closer),
	}

	s.updateTime()

	go func() {
		ticker := time.NewTicker(time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-closer:
				return
			case _, ok := <-ticker.C:
				if !ok {
					return
				}

				s.updateTime()
			}
		}
	}()

	return s
}

func (s *snowflake) Close() {
	s.closer.TrySend(struct{}{})
}

func (s *snowflake) updateTime() {
	seed := s.newSeed()
	old := s.seed.Load()
	for old < seed && !s.seed.CompareAndSwap(old, seed) {
		seed = s.newSeed()
		old = s.seed.Load()
	}

}

func (s *snowflake) newSeed() uint64 {
	return uint64(s.nodeId&0x1fff) |
		(uint64(time.Now().UnixMilli()<<(s.incrementLength+13)) & uint64(0x7f_ff_ff_ff_ff_ff_ff_ff))
}

func (s *snowflake) NewStr() string {
	return b62.Encode(s.New())
}

func (s *snowflake) New() uint64 {
	return s.seed.Add(incrementBase)
}
