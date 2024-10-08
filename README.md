# Allora Network

![AlloraLogo](assets/AlloraLogo.jpeg)

[![Go Report Card](https://goreportcard.com/badge/github.com/allora-network/allora-chain)](https://goreportcard.com/badge/github.com/allora-network/allora-chain)

![Docker!](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![Go!](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Apache License](https://img.shields.io/badge/Apache%20License-D22128?style=for-the-badge&logo=Apache&logoColor=white)

The Allora Network is a state-of-the-art protocol that uses decentralized AI and machine learning (ML) to build, extract, and deploy predictions among its participants. It offers actors who wish to use AI predictions a formalized way to obtain the output of state-of-the-art ML models on-chain and to pay the operators of AI/ML nodes who create these predictions. That way, Allora bridges the information gap between data owners, data processors, AI/ML predictors, market analysts, and the end-users or consumers who have the means to execute on these insights.

The AI/ML agents within the Allora Network use their data and algorithms to broadcast their predictions across a peer-to-peer network, and they ingest these predictions to assess the predictions from all other agents. The network consensus mechanism combines these predictions and assessments, and distributes rewards to the agents according to the quality of their predictions and assessments. This carefully designed incentive mechanism enables Allora to continually learn and improve, adjusting to the market as it evolves.

## Documentation

For the latest documentation, please go to <https://docs.allora.network/>

nano docker-compose.yaml

docker build -t alloranetwork/allora-chain:v0.5.0-docker-upgrade .

docker compose up -d

docker compose logs -f

docker exec -it allora_validator /bin/bash

allorad keys list
