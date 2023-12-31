# 1. Abstract

In the digital age characterized by a relentless pursuit of convenience and adaptability, the RentWave project is poised to revolutionize the rental services industry by seamlessly integrating Microservices (MS) and Event-Driven Architecture (EDA). This visionary approach involves breaking down complex monolithic applications into agile, independently deployable services while enhancing real-time responsiveness and system flexibility. RentWave adeptly tackles the challenges of development intricacies, real-time interactions, scalability, and operational reliability, weaving a highly agile, responsive, and scalable platform for users seeking a wide range of rentals, from books to vehicles and guest rooms. The project's comprehensive full-stack web application employs unique technology stacks for each microservice, implements EDA for seamless communication, leverages cloud solutions for hosting, and establishes a robust deployment pipeline through DevOps technologies. As RentWave unfolds, it is poised to further enrich offerings and redefine the boundaries of convenience and adaptability in the world of rental services.

# 2. Contents

1. [Abstract](#1.-abstract)
2. [Contents](#2-contents)
3. [Introduction](#3-introduction)
    1. [What are MicroServices](#31-what-are-microservices)
    2. [Approaches To Benefit From MicroServices](#32-approaches-to-benefit-from-microservices)
    3. [Scope And Objectives](#33-scope-and-objectives)
4. [Background](#4-background)
    1. [Literature Review](#41-literature-review)
    2. [Successful Implementations](#42-successful-implementations)
    3. [Failed Implementations](#43-failed-implementations)
    4. [Event Driven Systems](#44-event-driven-systems) 
    5. [Summary of Challenges](#45-summary-of-challenges)
5. [RentWave Overview](#5-rentwave-overview)
    1. [Key Challenges](#51-key-challenges)
6. [Design Review](#6-design)
    1. [Requirements](#61-requirements)
    2. [Overview Of Microservices](#62-overview-of-microservices)
    3. [Detailed Design](#63-detailed-design)
7. [RentWave Implementation](#7-rentwave-implementation)
8. [Summary](#8-summary)
9. [Appendix A - Getting Started With Microservices](#9-appendix-a---getting-started-with-microservices)
10. [Appendix B - RentWave](#10-appendix-b---rentwave)
11. [References](#11-references)

# 3. Introduction

## 3.1 What are MicroServices

Microservices is an architecture pattern used in designing and developing software applications in which an application is built as small, loosely coupled and independently deployable services. This is an alternative to the monolithic pattern where all services are tightly coupled. A microservices based system is broken down into a set of individual services where each service is responsible for a specifric and well-defined functions. These microservices generally expose APIs which are used to establish communication between them while also remaining loosely coupled.

## 3.2 Approaches To Benefit From MicroServices

Microservices can prove to be extremely beneficial in many aspects like development time, scalability, reliability to name a few of them. A bad design however, can be equally devastating to scalability and costs of an application. There are many approcahes to design a good microservices based system. 

Some of the most important ones are:

- **Clear Service Boundaries**: Define clear boundaries for each microservice, ensuring that they have specific, well-defined responsibilities. Avoid overlapping functionality between services.

- **Independent Development and Deployment**: Develop and deploy each microservice independently. This allows for faster development cycles and reduces downtime during updates.

## 3.3 Scope And Objectives

This project aims to explore the various methodologies available to design and develop a Micro-services based application called **RentWave** while simultaneously evaluating the pros and cons in various aspects of design and development which serves to emulate a real-world software application which needs high level of scalability, security, reliability and cost-effectiveness.

RentWave aims to have the following objectives accomplished:

- Designing and Development of various microservices

- Seamless interaction between the various services through APIs

- Use event driven architecture to further decouple services

- Deploy the project online and ensure scalability

- Develop a deployment pipeline using DevOps Technologies like Docker and GitHub Actions.

These objectives will be used as the evaluation metric during the design process of the application.

# 4. Background

## 4.1 Literature Review

Google defines a microservices architecture as an application that is developed as individual services. Each individual service is designed to handle a specific functionality of the application that is as independent as possible from other parts of the application. In case a service needs to communicate with other services, an interface can be exposed by each service so that each request can be parsed by a service as per its own requirements. It also suggests to containerize all the microservices as it simplfies the development by allowing a developer to focus on the functionalities rather than fixing various dependencies.[[1]](#what-is-microservices-architecture--google-cloud)

Scalability works differently in a microservices based design than that of a monolithic application. Scaling a monolithic application is a comparatively simpler process where allocating more resources to the existing system can work when the traffic on the application increases. However, scaling in a monolithic application is a rigid process where independent scaling of services is not feasible since the entire application needs to be scaled even if the objective is to scale up a single service in the system. This is solved by using the microservices pattern where each service can be scaled independently of each other which results in a more cost-effective scaling pipeline.[[2]](#microservice-scalability-challenges-and-how-to-overcome-them---developercom)

Microservices however also introduce new problems like "dependency hell" which arise in cases where multiple microservices share common dependencies like a library or a package. If a change in the library is needed to introduce a new feature or modify an existing functionality this might cause breaking changes in the other services. Not upgrading the library in the other services might work for a while but eventually the two different versions of the same library go out of sync which causes any further changes even remotely dependent on the library to become impossible. There also can arise circular dependencies where say a microservice X relies on Y and Y relies on Z but Z relies on X which makes any change in these services to require a lot of effort so as to not introduce bugs or worse, downtime which can cost billions in large corporations. [[3]](https://www.linkedin.com/pulse/dependency-hell-microservices-how-avoid-nabil-hijazi)

Event-Driven Architecture (EDA) is a design concept that helps in decoupling services by emitting events which can then be listened to by the other services. This means that the service emitting the service need not know anything about the service that would be consuming the event which would result in a more reliable system even if one of the instances of a microservice fails, since the event that was to be handled by thsat instance can now be handled by the other instances or even delay handling the event until the instance comes back online. EDA also enables easier scaling since it simplifies the addition of new event producers and the load balancing of the event consumers is automated where the events are redistrubted to the working nodes in case of a failure in one of them. [[4]](https://blog.hubspot.com/website/event-driven-architecture#What-is-event-driven-architecture)

## 4.2 Successful Implementations

### Amazon

Amazon is one of the first of the large corporations today to have successfully implemented a microservices pattern with over a few hundred microservices that are currently in use. Amazon started off with a monolithic system in the early 2000's but as the number of developers and traffic on their applications increased development in their system proved to be a challenge as well as scaling up their monolith. As it adopted the microservices way of developing their applications, they soon realised massive improvements in performance and ease in scalability.[[3]](#why-and-how-netflix-amazon-and-uber-migrated-to-microservices-learn-from-their-experience)


During the migration each individual functionality was identified and the highly inter-dependent code was separated into small chunks of code that provided with a small and specific function. Once all the individual chunks were identified and separated, they were put together into sections where dependent parts were put together in one section and the ones that are independent were put in other sections. This meant that the code in one section was indepedent of code in the other section so that the developers can work parallely in different sections without introducing breaking changes to the system. In cases where some amount of dependency was unavoidable, the sections introduced interfaces which are the only means through which the services could communicate. This was the beginning of the Service Oriented Architecture that we are familiar with today.[[4]](#what-led-amazon-to-its-own-microservices-architecture)


Amazon also started developing the various solutions in regards with the microservices pattern of design like Amazon Web Services (AWS) which boasts over 200 fully working services in regards to various domains of development like data storage, logging services, cloud applications, etc.

### Netflix

Netflix is another of the instances where migrating to microservices impacted its success today in a major way. Netflix originally started off in 2007 as a DVD distributor when online streaming was not a thing yet. By 2008, it was struggling to keep its services up and a major turning point came when its services went down for three days straight due to a database corruption. This led to them opting to a horizontally scalable database system rather than their existing vertically scaled database where they had no backup and an outage in it could have disastrous outcome.[[5](#completing-the-netflix-cloud-migration)]

In 2009, Netflix began designing its microservices architecture and chose AWS as its cloud provider. This migration to cloud also enabled to extend its services to a wider audience as well as improving its reliability and uptime. Netflix claims that the migration to cloud improved its uptime to **four nines**.
Netflix hit two billion API requests per day to its API Gateway in 2013 and by 2017 it had over 700 microservices functioning on the cloud. Today, Netflix boasts over a 1000 microservices and over 40 milllion users worldwide yet functions smoothly due to its transition to microservices.[[6]](#optimizing-the-netflix-api)


## 4.3 Failed Implementations

### Prime Video Quality Analysis [[7]](#scaling-up-the-prime-video-audiovideo-monitoring-service-and-reducing-costs-by-90)

Amazon's Prime Video is a massive streaming platform worldwide with over 200 million users. The Video Quality Analysis (VQA) team at prime video designed a monitoring tool for liverstreams but was not originally designed for high scalalbility was aimed to monitor a few thousand concurrent livestreams and scale it up gradually over time. However, it realised that the system was not performing well when it started adding more streams to its platform and also noticed the high costs of scaling up its infrastructure to the required standards.

They initally came up with a distributed microservice architecture that mainly consisted of three components i.e, Media converter to convert the stream into a buffer that is then to a Defect Detector to identify any issues with the stream and finally an orchestration service to handle the orchestration and control the flow. This design was theoretically scalable but they hit a hard limit of 5% of the scale they needed to achieve and the cost of scaling was not good either.
The reason for this was mainly due to the costs of passing the video data around the microservices which they tried to resolve by storing them temporarily in a S3 bucket but the calls to the bucket were getting expensive.

They realised that the microservices pattern was not giving them much benefits over the design they had earlier and switched to a monolith pattern. To resolve the scalability issue they parametrized their service so that they could specify each service to handle a different subset of the data. This migration to a monolith architecture from microservices reduced their costs by 90% as well as increasing scalability of their monitoring tool.

### Twitter And Wix

Most organisations do massively benefit from a microservice architecture but this also over complicates their system by adding dependencies that are not essential to the functioning of their applications.

Twitter is a massive social media platform with millions of users worldwide which also provides a lot of advertising services as well provides APIs to developers to develop applications making use of it. Twitter's CEO Elon Musk claims that only a mere 20% of microservices in twitter are essential for its functioning while the rest provide little to no functionality or performance improvements.[[8]](#elon-musk-on-twitter)

Wix is another such instance where migrating to microservices proved to be difficult. Wix began its migration to microservices from its monolithic system when they discovered that bugs in one part of the system had the potential to bring down the entire application. They encountered new issues like addressing failures and debugging which were not the case in their monolith system. They had create an entire framework to work with microservices. The employees at Wix also had to face challenges to migrate to microservices from monolith which they were used to which resulted in a culture shift. [[9]](#wixcom-migration-to-microservices)

## 4.4 Event Driven Systems

### Types Of Events

#### RESTful Events

These events contain data in the payload that usually represents the data that is received from RESTful APIs. These events are usually used to transfer data to the end users instead of the internal system like other microservices. These events however require additional consideration while designing since an event that is meant for a user is supposed to reach that user only which makes using events less secure for sensitive information.

#### Signal Events

These events contain tiny payloads and are used to "signal" the occurence of an event in a part of the system. These usually only contain the ID(s) of the object that is affected by the event. These events are also highly secure owing to the fact that it exposes very little to no information to the other parts of the system. However, the tiny payload also causes increased latency where more information about an object is needed since it requires additional API calls to the services.

#### Domain Events

These events are used to capture business-relevant events like `OrderPlaced`, `OrderUpdated`, etc. These events are generally more verbose than signal events and thus are also used as a history record. These events can also be used in CQRS patterns in databases where a write event is published when a change occurs in the database which is then consumed by the query counterpart of the database and updates the data accordingly.

## 4.5 Summary Of Challenges

### Scalability

Most organizations do migrate to microservices in their objective to become scalable but do not have a clear estimate of how high they might need to go. This results in either of the two cases:

- **Underestimating Scale**: The upper limit of the scaling of the application is massivley underestimated which though seems to work fine would result in scaling hard limits with the existing architecture. This can only be resolved by overhauling entire design process which would take a considerable amount of time as well as performance issues during that time due to the high load on the under powered system.

- **Overestimating Scale**: The converse is also troublesome where the scale severly overestimated and resources are allocated to the system which remain unused resulting in un-reasonable costs of infrastructure while providing absolutely no benefit.

### Communication Between Microservices

Microservices often need to communicate in some form or the other for various needs like authentication, fetching dependent data etc. This would mean that data needs to travel through long chains just to accomplish some basic functionality. This would also result in increased latency, lower throughput and higher network costs.

The long chains of communication also can introduce multiple points of failure which would mean that debugging and fixing issues in critical times would be time consuming and can result in massive losses in revenue. In case a particular service is down, the entire application has the potential to go down with it if the communication between microservices are not designed properly.

### Infrastructure and Design

The microservices need to be designed with careful thought into each and every aspect to avoid any unforeseen troubles and the infrastructure must be scoped out to ensure that the suitable infrastructure is available and within the budgets as well as how reliable the infrastructure is. In cases where the infrastructure itself fails, the application would go down and would be difficult to get it back up without finding suitable replacements for it which generally takes a considerable amount of time.

# 5. RentWave Overview

In the current age where convenience and adaptability are highly sought after, the RentWave project presents an innovative platform that aims to redefine the rental services landscape. RentWave acts as a comprehensive hub for a wide range of commodities, providing a versatile solution for individuals looking to rent various items, including books, vehicles, and guest rooms. The project's primary goal revolves around developing a robust, full-stack web application built on microservices principles, each utilizing distinct technology stacks to mirror real-world application development.

The central aspect of this endeavor is the adoption of an event-driven architecture, enhancing dynamism and responsiveness. By combining this architectural approach with cutting-edge cloud solutions and the implementation of DevOps best practices, RentWave emerges as a frontrunner in modern software design. This integration of advanced technologies and forward-thinking methodologies positions RentWave as an industry influencer, setting new standards for the seamless and adaptable rental experience that today's consumers expect. In an era where convenience and adaptability are key, RentWave signifies the innovative progress towards a future where renting various items becomes more accessible, efficient, and hassle-free.

## 5.1 Key Challenges Addressed

### Data Ownership In Microservices

Firstly, dealing with the complexity of building the platform is a big challenge. Making a platform that can handle many types of rentals and user features can be really hard. To tackle this, RentWave uses a microservices approach. Each microservice focuses on a specific part of the platform, which makes it easier to develop and maintain, making the whole process more straightforward and understandable.

### Event Driven Architecture

In addition, the imperative need for real-time interaction demands a strategically engineered solution. Users now anticipate instantaneous engagement and up-to-the-minute updates. In response, RentWave strategically employs an event-driven architecture. This architectural decision substantially addresses the task of establishing smooth communication among the diverse components of the platform, thereby ensuring users remain seamlessly connected and continuously informed.

Event-driven architecture bestows several technical advantages upon RentWave. It fosters an environment where components of the platform can respond to events, such as user actions or external triggers, in real-time. This dynamic responsiveness enables immediate updates and feedback to users, enhancing their experience. Moreover, it facilitates the decoupling of various system elements, allowing them to function independently, which not only simplifies development and maintenance but also enhances system resilience.

By utilizing event-driven architecture, RentWave can also scale horizontally to accommodate growing user demands, as each microservice can handle events and workloads independently. This scalability ensures that the platform can seamlessly expand its capacity to meet the demands of a growing user base, without compromising the quality of real-time interactions.

### Development And Deployment Pipeline

The aspect of Deployment and Operations within the RentWave project is a critical element in ensuring the platform's continuous functionality and improvement. To meet this imperative requirement, we harnessed the power of DevOps technologies, specifically leveraging tools such as Docker and GitHub Actions. This strategic adoption of DevOps principles serves not only to streamline the deployment process but also to orchestrate a seamless and efficient operational environment for the platform.

A pivotal facet of this endeavor involves the establishment of a robust Continuous Integration and Continuous Deployment (CI/CD) pipeline, which is indispensable in the modern software development landscape. This CI/CD pipeline plays a multifaceted role in RentWave's development ecosystem. It automates the integration of code changes, ensuring that new features, bug fixes, and enhancements are smoothly incorporated into the platform without disruptions or errors.

# 6. Design

## 6.1 Requirements

RentWave is a web-application that serves to simplify the renting process of various commodities like unused books and vehicles. To serve this purpose the following are the requirements that are planned to be accomplished by the application:

### Functional Requirements

- Users can register and login to the application to access the core functionalities.
- Users can post an Advertisement for their unused books and vehicles with a rental price they are happy with.
- Users can view and search for other users' advertisements so that they can rent items that they like.
- Users can rent the items they like for a duration that permissible by the owner of the advertisement.
- Users must be charged from their wallet automatically if they exceed the agreed upon duration.
- Users must be notified when their item has been bought for rent.

### Non-Functional Requirements

- Application must be highly scalable.
- Application must be secure and support for future external APIs.
- Application must be reliable and resisitant to independent failures of microservices.
- Depolyment must be quick and feasible.
- Maintain logs for each event for quick and easy resolution of issues.

## 6.2 Overview Of Microservices

RentWave makes use of several microservices in order to fulfill its requirements of authentication, notifications, scalability, etc.

The brief overview of the core microservices used is shown below:

![Microservices Overview](https://github.com/Animesh-03/TBD-Project/blob/master/docs/images/MicroServices%20Architecture.png?raw=true)

The functionalities offered by each of the above mentioned services will be described in detail in the following section.

## 6.3 Detailed Design

### User Service

This service stores the data related to the user like his address, name, etc. as well as acts as the authentication service by handling the registration and login of users. Upon login, the user is granted an access token that is then used to authenticate the user in the future requests. If the token turns out to be invalid or does not exist then the request fails and the user is denied access to any of the core functionalities of the application.

Below is a brief overview of how the authentication works:

![Authentication Flow](https://github.com/Animesh-03/TBD-Project/blob/master/docs/images/Authentication%20Flow.png?raw=true)

### Book Service

This service handles the users that add new books for rent. This service also handles toggling the availability of books for instances where the owner of the book does not want the book to be rented out. Each book is identified by a unique isbn number. If a user is adding a book whose isbn already exists then that advertisement is linked to that book to make it easier to query multiple advertisements of the same book since the isbn is the same.

This service emits event like `book.add`, `book.update` etc., which are used by other services like logger service to log when the book was added/modified or the notification service to notify users that a new book is available for rent.

### Vehicle Service

This service handles the users that add new vehicles for rent. This service also handles toggling the availability of vehicles for instances where the owner of the vehicle does not want the vehicle to be rented out. Each vehicle is tagged with a type of the vehicle so that the users can query a vehicle specific to their needs like size, seating, boot space, etc.

This service emits event like `vehicle.add`, `vehicle.update` etc., which are used by other services like logger service to log when the vehicle was added/modified or the notification service to notify users that a new vehicle is available for rent.

### Rental Service

The rental service manages the renting process by users who wish to rent out items from the avaialble pool advertisements. This service is used to display the current status of rented items of a user as well as it enforces a penalty on a user if they are late on returning the item.

The rental service will wait for a `rent.complete` event and if the event does not arrive before the agreed upon period of rent expires then it emits a `rent.late` event which will be used to deduct balance automatically from the wallet of the user.

### Kafka Cluster

Kafka Cluster is used as the event broker and serves a central role in the event driven system of RentWave. All the other microservices subscribe to events in the kafka cluster which then sends the events as soon it as receives to the subscribers. There are multiple broker nodes in the cluster to ensure reliability of the events and to make sure that events are not lost completely when a broker node goes down.

### Logger Service

The logger service serves only a single purpose i.e, to log every event that occurs in the application. The logger service is important as this provides timestamped events and logs which help in not only debugging but also to resolve issues that users face while using the application which is an important aspect of user experience in any application.

### API Gateway

The API Gateway is an important aspect of RentWave since it establishes communication between the various microservices in the application. The presence of gateway eliminates the need of every service to know the location of every other service since the gateway is the only service that knows the locations of the other services and all communication is channeled through the gateway.

# 7. RentWave Implementation

RentWave is designed such that the data need not move around services too much which reduces latency, network costs as well as improves security by not exposing data that is not required to accomplish a particular task. This is enabled by an efficient database design.

The ER Diagram below shows the data model in brief:

![ER Diagram](https://github.com/Animesh-03/TBD-Project/blob/master/docs/images/ER%20Diagram.png?raw=true)

## User Service and Book Service

The User Service and Book Service is implemented using golang[[10]](#golang) which is an extremely performant language and is widely used in various backend related services in the software development industry due to its wide support for various tools and libraries.

Peformance is vital in the user service since this is responsible for authentication which becomes extremely important and has high probability to have increased traffic where time of execution becomes the most important factor and also helps in reducing costs.

Golang is also a typesafe language which improves the security of the authentication service and prevents misuse of the APIs. The emerging popularity of the language also ensures that the service will remain maintainable in the future which is a serious problem for companies that use legacy code. The new developers do not have the expertise in the tech stack and this results in an unmaintanable codebase where no issues can be fixed.

Golang has the Gin[[16]](#gin-framework) framework which is used to build the server of these services and maps the handler function to the requests. Golang also provides the GO-ORM[[17]](#gorm) which serves as the ORM(Object Relations Mapper) for the server and the PostgreSQL database.

A PostgreSQL database is used since the schema in these services is relatively rigid and does not need to change throughout the iterations of development process.
This also helps in vertical scalability of these services since SQL databases are highly vertically scalable as compared to its NoSQL counterparts like MongoDB which are horizontally scalable.

### Authentication

The authentication in RentWave is an access token based implementation. Once a user enters the username and password to register, the password is then hashed using SHA-256[[12]](#sha-256) which is a popular asymmetric hashing algorithm[[11]](#asymmetric-hashing-algorithms) and then sent over to the gateway so that the gateway never receives the actual password itself as well as the password is never sent over the network to prevent MITM[[13]](#mitm-attack) attacks. The database stores the username and the hash of the password.

During login, after the user enters the username and password and presses the login button, the password is hashed and then sent to the gateway which then forwards the request to the user service. The user service fetches the hash of the password that it stored in the database during registration and compares that hash to the one that is provided by the user. If the hashes match then a JWT[[14]](#jwt) is generated where the payload has the user id and timestamp. This token is returned to the gateway which then converts the access token to a cookie by adding an expiration time and sent to the client.

This cookie is then sent along with every other request and is verified before the access to the actual API is given.

## API Gateway And Vehicle Service

The API Gateway and Vehicle Service are written in Typescript[[15]](#typescript) which is a typesafe version of javascript and helps improve development experience as well as time. This is in conjunction with NodeJS which is runtime environment for Javascript and enables running javascript outside of the context of a browser.

Typescript is chosen as the language here because of its compatibility with JSON. Most of the communication between the frontend and even between the microservices happens through HTTP APIs having JSON body as the payload. This enables the gateway to handle the requests from the frontend more efficiently. Since the gateway itself does not handle any business logic, the performance tradeoff is useful in favor of reliability using JSON.

Though the API Gateway itself is stateless and thus does not need a database, the vehicle service uses MongoDB for its storage needs as opposed to an SQL database since the schema for the service is likely to change to due to several factors like change in license policy or even general renting policies. A NoSQL database supports a changing schema more readily than the SQL variants.

## Rental Service

The rental service is the heart of the application which in addition to emitting the core events of the event driven system, it also ensures fast response times to requests as this is estimated to be the most used microservice in the application.

The microservice takes in the type of entity that is being rented and the entity ID along with user ID to interact with the renting aspect of the application. This handles requests pertaining to a user renting an entity, viewing active rentals, past history, etc.

The service emits events when a new entity is being rented which can then be consumed by the corresponding service to process the rental. This service also enforces the rental duration by emittiing events once the agreed upon duration of the rent has passed which can then be processed by deducting currency from the user's wallet automatically.

## Frontend

The frontend for the web application is written using SvelteKit[[18]](#sveltekit) which is written on top of Svelte[[19]](#svelte), a UI Framework. This is chosen as opposed to the more popular frameworks like NextJS which is built on top of ReactJS and AngularJS due to the following reasons:

- **Performance**: SvelteKit does not work on a virtual DOM like ReactJS or AngularJS and then update the acutal DOM which is significantly slower than working on actual DOM.

- **SEO**: ReactJS works by sending an empty HTML document in its intial request along with some JS which eventually populates the HTML with the content, this is generally fine but for a search engine which indexes pages based on their content. This is disastrous for ReactJS applications since the search engine indexes an empty HTML page with no content as it does not wait for the javascript to do its work.

- **Shifting Trend**: Though ReactJS is still hugely popular, it has been coming under attack due to various reasons which is causing a shift to other frameworks with Svelte being at the forefront of them. This ensures that the tech stack stays relevant in the future and new developers can continue developing new features and maintaining it.

## Deployment And DevOps

DevOps practices emphasize collaboration, communication, and automation between development and operations teams, aiming to streamline the software development lifecycle. It also enables to look ahead on the feasibility of a solution in terms of scalability and testing.

 By automating manual and repetitive tasks, DevOps enhances overall efficiency. Automated testing, deployment, and monitoring reduce the likelihood of human errors, allowing teams to focus on higher-value tasks and innovation rather than time-consuming, error-prone manual processes.

 RentWave uses some of the widely used tools and strategies to ensure a smooth, secure and efficient deployment of the microservices like **nginx**, **docker**, **GitHub Actions**, etc.,

### NGINX

Nginx is a high-performance web server and reverse proxy server. It serves as a middleman between a user's web browser and the web server hosting a website. Known for its efficiency and speed, Nginx handles tasks like serving static content, managing SSL/TLS connections, and load balancing to distribute incoming web traffic across multiple servers. Its event-driven architecture allows it to handle a large number of simultaneous connections with low resource usage. Nginx is often used to enhance website performance, improve security, and ensure smooth delivery of web content, making it a popular choice for websites with high traffic loads.

RentWave utilises Nginx's capability to forward requests to services on the basis of the origin of the request i.e, the host. The various micro-services from above are hosted on different hosts of the form `*.rentwave.live`. Nginx also upgrades all the incoming http connections to https by forwarding the requests to the corresponding https protocol.

The SSL certificates were generated using `certbot` which is an open source tool which uses `let's encrypt`, a widely recognized certificate authority. These certificates are installed on all the domains the services are hosted on to enable HTTPS traffic.

### Docker

Docker is a platform that enables developers to build, package, and distribute applications in lightweight, portable containers. Containers encapsulate everything needed to run an application, including the code, runtime, libraries, and dependencies, making it consistent across different environments. This approach is particularly advantageous for microservices architecture.

Moreover, Docker provides a straightforward and consistent way to scale microservices horizontally. Horizontal scaling involves adding more instances of a microservice to the system to handle increased demand or distribute the workload effectively. With Docker, this process becomes seamless. Developers can easily replicate containers and deploy them across multiple servers or cloud instances. The ability to scale horizontally ensures that microservices can efficiently handle growing user loads or sudden spikes in demand without compromising performance.

In the context of RentWave, each microservice is designed to function independently wherever possible which makes scaling a particular service extremely easy. This can be done by starting the docker instances of all the dependencies of just that service instead of having to scale the entire application even if the load on the other services is minimal.

### Cloud Infrastructure

RentWave is deployed on the oracle cloud which enables leveraging various advantages of cloud. A few of the important factors are listed below:

- **Resource Efficiency**: Microservices often have varying resource requirements. Cloud infrastructure allows dynamic allocation and de-allocation of resources based on the specific needs of each microservice. This ensures efficient utilization of computing resources, minimizing costs and optimizing performance.

- **Fault Tolerance and High Availability**: Cloud providers typically operate in multiple data centers across different geographic regions. This inherent redundancy enhances fault tolerance and ensures high availability for microservices. If one data center experiences issues, traffic can be redirected to another, minimizing downtime.

- **Cost Efficiency**: Cloud services operate on a pay-as-you-go model, allowing organizations to pay only for the resources consumed. This cost-effective approach is well-suited for microservices, where individual components can be scaled independently based on demand, avoiding unnecessary expenses.

- **Global Reach**: Cloud providers have a global network infrastructure, enabling microservices to be deployed close to end-users for reduced latency. This is crucial for applications with a diverse user base, ensuring a consistent and responsive user experience.

# 8. Summary

Microservices simplify the development and maintenance of applications by breaking them into smaller, manageable services. Event-Driven Architecture enhances real-time responsiveness and flexibility. However, adopting these technologies comes with challenges, including development complexities, real-time interactions, scalability, and operational reliability.

The implementation of RentWave helped in realising the pros and cons of using a microservices architecture along with an event driven system and the challenges that arise when using various tech stacks for each microservice and maintaining consistent data between all the microservices.

## Future Work

There is a lot more scope in leveraging the strengths of event-driven architecture in the existing system to enable a more robust system and seamless handling of requests. A significant portion of getting started with a microservices system can be simplified by eliminating the need to write boilerplate code for various purposes like Dockerization, DevOps, service discovery and communication between them.

A tool/framework can be developed to make this part of the development process smoother and easier. This tool can be used to setup a basic microservices system by using something simple like a config file which contains the specifications of the various services needed along with the tech stack that each service would be based on. This however should be not hinder further development by black-boxing major aspects of the system and should remain highly configurable. This tool should also setup a basic deployment pipeline which is again configurable through the config file.

## Conclusion

In conclusion, the adoption of microservices in production environments brings forth a myriad of benefits, including enhanced scalability, fault isolation, and the flexibility to employ diverse technologies. This approach fosters agility and continuous deployment, allowing for faster and more efficient development cycles. However, it is crucial to acknowledge the challenges associated with this paradigm.

Efficient integration of DevOps practices is paramount for managing the complexity introduced by microservices. Additionally, handling data management and migration, orchestrating service communication using event-driven architecture, and ensuring effective service discovery and load balancing pose significant challenges. Security and governance, along with the intricacies of monitoring, tracing, and testing in a multi-service environment, require meticulous attention.

Successfully navigating these challenges demands a holistic approach, incorporating robust DevOps collaboration, meticulous data management strategies, and thoughtful implementation of event-driven communication. Moreover, addressing security concerns and establishing comprehensive monitoring and testing practices are essential components of a successful microservices implementation. In essence, while the benefits are substantial, embracing microservices in production necessitates a strategic and comprehensive approach to build resilient, scalable, and efficient systems.

# 9. Appendix A - Getting Started With Microservices

Microservices in simple terms are just multiple servers that communicate with each other using APIs and serve purposes that the end user needs. Each microservice can use their own tech stack and database but must expose APIs.

To demonstrate a simple microservice application, let us consider an API Gateway service whose sole purpose is to process requests from the users and convert the data recieved into the format that other services can consume using their own APIs. Let the other microservice in this case be a simple user authentication service using [JWTs](#jwt).

## User Auth Service

Let the User Authentication service use the following tech-stack:
- Golang
- GIN
- GORM
- Postgres

The detailed documentation and how to use them in the project can be found in the references.

To create a new Golang project run `go mod init user-service` and install the above dependencies.

Setup the connection between the database and the auth service using `GORM`:
```go
dsn := "host=postgres user=user password=pass dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
    fmt.Println("Could not connect to  db")
    return
}

// Close the connection at the end
connection, _ := db.DB()
defer connection.Close()
```

Now initialize the `gin` server by listening to a port:
```go
app := gin.Default()

app.POST("/register", RegisterHandler)
app.POST("/login", LoginHandler)

app.Run("0.0.0.0:8087")
```

Create the user model:
```go
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string
}
```
Create two functions i.e, `RegisterHandler` and `LoginHandler` which are called when the `register` and `login` endpoints are called respectively.

```go
func RegisterHandler(c *gin.Context) {
	var registerCredentials models.RegisterCredentials
	userRepository := repositories.NewUserRepository(global.DB)

	if err := c.BindJSON(&registerCredentials); err != nil {
		fmt.Println("Invalid Request Body: ", err)
		c.IndentedJSON(403, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	err := auth.Register(userRepository, registerCredentials.Username, registerCredentials.Password)
	if err != nil {
		fmt.Println("Error registering user: ", err)
		c.IndentedJSON(500, gin.H{
			"error": "Username already exists",
		})
		return
	}

	c.IndentedJSON(200, "OK")
}

func LoginHandler(c *gin.Context) {
	var loginCredentials models.LoginCredentials
	userRepository := repositories.NewUserRepository(global.DB)

	if err := c.BindJSON(&loginCredentials); err != nil {
		fmt.Println("Invalid Request Body: ", err)
		c.IndentedJSON(403, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	token, err := auth.Login(userRepository, loginCredentials.Username, loginCredentials.Password)
	if err != nil {
		c.IndentedJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(200, gin.H{
		"access_token": token,
	})
}
```

Note that the `auth` and `repository` packages are simple wrappers around `gorm`. To learn more about these packages feel free to visit the [GitHub Repository](https://github.com/Animesh-03/RentWave).

Finally dockerize the go application by creating a `Dockerfile`:
```Dockerfile
FROM golang:1.19

WORKDIR /usr/local/go/src/user

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
```

## API Gateway

Now create NodeJS application for the API Gateway that uses the following tech-stack:
- Typescript
- ExpressJS
- jsonwebtoken

To begin with initialize an express server by listening on a port:
```ts
import express from "express";

const app = express();

app.use(express.json());

app.get("/", (_, res) => {
    res.send("Hello World");
});

app.listen(8082, () => console.log("Started API Gateway on PORT 8082"));
```

Ideally the gateway would be handling requests to various services and thus to better organize the code, make use of express routers to modularize the endpoints and handlers. Create a new folder called `routers` and create a file called `auth.ts`.
```ts
// auth.ts

import { Router } from "express";
import axios, { AxiosError } from "axios";

const userRouter = Router();

userRouter.post("/register", async (req, res, next) => {
    // Redirect to User Service
    const registerResponse = await axios.post(`${process.env.USER_SERVICE_URL}/register`, {
        "username": req.body.username,
        "password": req.body.password
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            res.status(err.response?.status ?? 500).send(err.response?.data);
            next();
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
            next();
        }
    })


    if(registerResponse)
        res.sendStatus(200);
});

userRouter.post("/login", async (req, res, next) => {
    // Redirect to User Service
    const loginResponse = await axios.post(`${process.env.USER_SERVICE_URL}/login`, {
        "username": req.body.username,
        "password": req.body.password
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
        next();
    });

    if(loginResponse) {
        // If request is sucessful then set the auth cookie while sending response
        res.cookie("auth", loginResponse.data.access_token, {
            // Expire after 24hrs
            maxAge: 24*60*60*1000,
            sameSite: "none",
            secure: true
        });
        res.sendStatus(200);
    }

});
```

The user router handles the requests to the `register` and `login` endpoints and returns the response from the auth service which is then returned back to the user.

Now use the user router:
```ts
//app.ts

app.use("/user", userRouter);

// If we have another microservice to handle orders
app.use("/order", orderRouter);
```

Now dockerize the gateway by creating a Dockerfile:
```Dockerfile
FROM node:alpine

WORKDIR /gateway

COPY ./package.json ./
RUN yarn install

COPY . .

CMD ["yarn", "dev"]
```

## Microservices Docker Compose

Finally use docker compose to startup all the microservices and their dependencies:
```yml
version: '3.1'

services:

    gateway:
    build: ../gateway/
    ports:
        - "8082:8082"
    environment:
        USER_SERVICE_URL: "http://user-service:8087"
    volumes:
        - ../gateway/src:/gateway/src

    postgres:
        image: postgres
        restart: always
        ports:
        - "5432:5432"
        environment:
        POSTGRES_USER: user
        POSTGRES_PASSWORD: pass
        POSTGRES_DB: users

    user-service:
        build: ../user-service/
        ports:
        - "8087:8087"
```

The microservices application can be run by the command `docker compose up --build`.

# 10. Appendix B - RentWave

RentWave is an application that assists users in the rental lifecycle of products like books and vehicles.

## Login and Register Pages
Users can register and login to RentWave.

![Login Page](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Login.png?raw=true)

![Register Page](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Register.png?raw=true)

## Dashboard

Users can visit the dashboard after logging where they can view the currently available rentals.

![Dashboard](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Dashboard.png?raw=true)

## Search and Filters

Users can also search and filter for products they are interested in.

![Search](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Search.png?raw=true)

![Filter](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Filter.png?raw=true)

## Adding An Item For Rent

Users can add products to be available for rent by going entering the required details.

![Add For Rent](https://github.com/Animesh-03/RentWave/blob/master/docs/images/Add%20Book.png?raw=true)


# 11. References

1. ##### [What is Microservices Architecture?- Google Cloud](https://cloud.google.com/learn/what-is-microservices-architecture)
2. ##### [Microservice Scalability Challenges and How to Overcome Them - developer.com](https://www.developer.com/design/microservices-scalability-challenges/)
3. ##### [Why and How Netflix, Amazon, and Uber Migrated to Microservices: Learn from Their Experience](https://www.hys-enterprise.com/blog/why-and-how-netflix-amazon-and-uber-migrated-to-microservices-learn-from-their-experience/)
4. ##### [What Led Amazon to its Own Microservices Architecture](https://thenewstack.io/led-amazon-microservices-architecture/)
5. ##### [Completing the Netflix Cloud Migration](https://about.netflix.com/en/news/completing-the-netflix-cloud-migration)
6. ##### [Optimizing the Netflix API](https://netflixtechblog.com/optimizing-the-netflix-api-5c9ac715cf19)
7. ##### [Scaling up the Prime Video audio/video monitoring service and reducing costs by 90%](https://www.primevideotech.com/video-streaming/scaling-up-the-prime-video-audio-video-monitoring-service-and-reducing-costs-by-90)
8. ##### [Elon Musk on Twitter](https://twitter.com/elonmusk/status/1592177471654604800)
9. ##### [Wix.com Migration to Microservices](https://techbeacon.com/app-dev-testing/monolith-microservices-horror-stories-best-practices)
10. ##### [Golang](https://go.dev)
11. ##### [Asymmetric Hashing Algorithms](https://www.cryptomathic.com/news-events/blog/differences-between-hash-functions-symmetric-asymmetric-algorithms)
12. ##### [SHA-256](https://security.stackexchange.com/questions/41431/reason-for-using-asymmetric-encryption-along-with-sha256)
13. ##### [MITM Attack](https://www.onelogin.com/learn/6-types-password-attacks)
14. ##### [JWT](https://jwt.io)
15. ##### [Typescript](https://www.typescriptlang.org/)
16. ##### [Gin Framework](https://gin-gonic.com/)
17. ##### [Gorm](https://gorm.io/index.html)
18. ##### [SvelteKit](https://kit.svelte.dev/)
19. ##### [Svelte](https://svelte.dev/)
20. ##### [Types Of Events](https://blog.frankdejonge.nl/the-different-types-of-events-in-event-driven-systems)