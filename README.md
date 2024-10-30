# Blockchain-Based Asset Management System with Hyperledger Fabric

This project develops a blockchain solution for managing and tracking assets through Hyperledger Fabric. It is organized into three main components:

## Project Summary

- **Level-1**: Establish a test network for Hyperledger Fabric.
- **Level-2**: Create and validate a smart contract to handle asset operations.
- **Level-3**: Develop a REST API to interact with the deployed smart contract.

## Directory Layout

- `fabric-test-network/`: Contains the files and configurations for setting up the Hyperledger Fabric test network.
- `smart-contract/`: Implements the smart contract for asset management, developed in Go.
- `rest-api/`: Provides the REST API for interacting with the smart contract, including Docker setup for easy deployment.

## Getting Started

### Requirements

To get started, ensure the following software is installed:

- Docker
- Golang
- Hyperledger Fabric binaries
- Node.js (necessary for setting up the network)

### Level-1: Setting Up the Test Network

1. Follow the instructions in the [Hyperledger Fabric Test Network Documentation](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html) to set up the network environment.

### Level-2: Smart Contract Creation

1. Refer to the [Hyperledger Fabric Chaincode Guide](https://hyperledger-fabric.readthedocs.io/en/latest/smartcontract/smartcontract.html) for guidance on developing and testing the smart contract, which will facilitate asset creation, updating values, and transaction history management.

### Level-3: REST API Development

1. Build a REST API in Go that enables interaction with the smart contract, allowing clients to create, read, update, and delete assets.
2. Package the API as a Docker image for streamlined deployment and scalability.

## Useful Resources

- [Hyperledger Fabric Documentation](https://hyperledger-fabric.readthedocs.io/en/latest/)
- [Hyperledger Fabric Samples Repository](https://github.com/hyperledger/fabric-samples)

