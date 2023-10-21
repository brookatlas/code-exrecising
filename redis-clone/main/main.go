package main

import redisclone "redis-clone"

func main() {
	redis_clone_client := redisclone.RedisCloneClient{
		Host: "localhost",
		Port: 6379,
	}

	redis_clone_client.Run()
}
