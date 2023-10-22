package main

import redisclone "redis-clone"

func main() {
	redis_clone_client := redisclone.NewRedisCloneClient(
		"localhost",
		6379,
	)

	redis_clone_client.Run()
}
