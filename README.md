# EFISHERY MONOREPO

efishery monorepo is made for efhisery technical test purposes. I'm using python to handle authentication. The token generated from this python will be used later in the fetch-app that uses golang. Every request received by fetch-app will check the JWT token to auth-app. If jwt is invalid it will return an error. 

In addition, the fetch-app also checks the roles and permissions of each user. For more details, you can see the following diagram. After that, proceed with fetching data to get commodity and currency data from an external API.In order not to burden the server, we do caching when making requests to external APIs
## auth-app
Microservice created for registration, login, and token checking. Built with the python programming language and more emphasis on a simple structure. You can check more information regarding this API at the following [link](https://github.com/hudabikhoir/efishery-monorepo/blob/master/auth-app/README.md)

## fetch-app
Microservice build with golang and hexagonal architecture to fetch commodities data. I use Linq to make query in struct data golang. The API architecture that we created emphasizes a neater code structure and modules because we adopt a hexagonal architecture. In this architecture already supports several database drivers such as mariadb, mongodb and sqlite. You can check more information regarding this API at the following [link](https://github.com/hudabikhoir/efishery-monorepo/blob/master/fetch-app/README.md)

## api documentation
All features of the two APIs are well documented in the following link [link](https://github.com/hudabikhoir/efishery-monorepo/blob/master/API.md)
