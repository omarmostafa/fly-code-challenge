
# Fly Code Challenge

The main idea of the task is to build an API that read the payment transaction data from files (Json format) and return them in the API response as json format. There are two payment providers flypayA and flypayB.

## Install 
Clone project 

```bash
git clone https://github.com/omarmostafa/fly-code-challenge.git
``` 

## Requirements

Install docker compose 

## Run App and test cases

 ```bash
docker-compose up --build
 ``` 

You will see the unit test result while build the project

 ## Run API 
  ```
 localhost:8080/api/payments/transactions
  ``` 
 
 ## Intro
 
 Implement this API `/api/payment/transaction `
 -  list all transaction which combine transactaions from all the available provider(`flypayA` and `flypay`B)
 -  filter transaction by payment providers for example `/api/payment/transaction?provider=flypayA` it should return transaction for flypay only
 -  filter transaction three statusCode (`authorised`, `decline`, `refunded`) for example /api/payment/transaction?statusCode=authorised it should return all the transactions from all payment providers that have status code authorised
 -  filer by amount range for example `/api/payment/transaction?amountMin=10&amountMax=100` it should return all transaction between 10 and 100 including 10 and 100
 -  filer by `currency` 
 -  combine all this filter together 

 
 ## Entities 
 
 Entities is the data object for our providers , transactions and filter
 
 ## Repositories 
 
 Repositories is responsible to get the data from jsons file and retrieve it 
 
 ## Mappers
 
 Its responsible to get the data from repo and map it to transaction entity with the common schema
 

## Handlers

Its responsible for Manage filters and providers to filter transactions and get the right provider

## Scalability
- if we need to add more providers all we need is create an entity with it's schema and create its repository to get the data from the right file and create mapper to map it to the common schema

- if we need to add more filters all we need it to add it to filter entity and add its criteria in filter handler