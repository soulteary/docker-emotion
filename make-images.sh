docker build -t soulteary/bert-base-uncased-emotion:2022.09.30 -f docker/bert-base-uncased-emotion/Dockerfile .

docker build -t soulteary/novel-zh-en:2022.09.30 -f docker/novel-zh-en/Dockerfile .

docker build -t soulteary/emotion:base-2022.09.30 -f docker/web/Dockerfile.base .

docker build -t soulteary/emotion:2022.09.30 -f docker/web/Dockerfile .
