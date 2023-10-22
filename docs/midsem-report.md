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
    4. [Summary of Challenges](#44-summary-of-challenges)
5. [RentWave Overview](#5-rentwave-overview)
    1. [Key Challenges](#51-key-challenges)
6. [Design Review](#6-design-review)
    1. [Requirements](#61-requirements)
7. [RentWave Implementation](#7-rentwave-implementation)
8. [Conclusion](#8-conclusion)
9. [Appendix](#9-appendix)
10. [References](#10-references)

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

Amazon is one of the first of the large corporations today to have successfully implemented a microservices pattern with over a few hundred microservices that are currently in use. Amazon started off with a monolithic system in the early 2000's but as the number of developers and traffic on their applications increased development in their system proved to be a challenge as well as scaling up their monolith. As it adopted the microservices way of developing their applications, they soon realised massive improvements in performance and ease in scalability.

During the migration each individual functionality was identified and the highly inter-dependent code was separated into small chunks of code that provided with a small and specific function. Once all the individual chunks were identified and separated, they were put together into sections where dependent parts were put together in one section and the ones that are independent were put in other sections. This meant that the code in one section was indepedent of code in the other section so that the developers can work parallely in different sections without introducing breaking changes to the system. In cases where some amount of dependency was unavoidable, the sections introduced interfaces which are the only means through which the services could communicate. This was the beginning of the Service Oriented Architecture that we are familiar with today. 

Amazon also started developing the various solutions in regards with the microservices pattern of design like Amazon Web Services (AWS) which boasts over 200 fully working services in regards to various domains of development like data storage, logging services, cloud applications, etc.

### Netflix

Netflix is another of the instances where migrating to microservices impacted its success today in a major way. Netflix originally started off in 2007 as a DVD distributor when online streaming was not a thing yet. By 2008, it was struggling to keep its services up and a major turning point came when its services went down for three days straight due to a database corruption. This led to them opting to a horizontally scalable database system rather than their existing vertically scaled database where they had no backup and an outage in it could have disastrous outcome.

In 2009, Netflix began designing its microservices architecture and chose AWS as its cloud provider. This migration to cloud also enabled to extend its services to a wider audience as well as improving its reliability and uptime. Netflix claims that the migration to cloud improved its uptime to **four nines**.
Netflix hit two billion API requests per day to its API Gateway in 2013 and by 2017 it had over 700 microservices functioning on the cloud. Today, Netflix boasts over a 1000 microservices and over 40 milllion users worldwide yet functions smoothly due to its transition to microservices.


## 4.3 Failed Implementations

### Prime Video Quality Analysis

Amazon's Prime Video is a massive streaming platform worldwide with over 200 million users. The Video Quality Analysis (VQA) team at prime video designed a monitoring tool for liverstreams but was not originally designed for high scalalbility was aimed to monitor a few thousand concurrent livestreams and scale it up gradually over time. However, it realised that the system was not performing well when it started adding more streams to its platform and also noticed the high costs of scaling up its infrastructure to the required standards.

They initally came up with a distributed microservice architecture that mainly consisted of three components i.e, Media converter to convert the stream into a buffer that is then to a Defect Detector to identify any issues with the stream and finally an orchestration service to handle the orchestration and control the flow. This design was theoretically scalable but they hit a hard limit of 5% of the scale they needed to achieve and the cost of scaling was not good either.
The reason for this was mainly due to the costs of passing the video data around the microservices which they tried to resolve by storing them temporarily in a S3 bucket but the calls to the bucket were getting expensive.

They realised that the microservices pattern was not giving them much benefits over the design they had earlier and switched to a monolith pattern. To resolve the scalability issue they parametrized their service so that they could specify each service to handle a different subset of the data. This migration to a monolith architecture from microservices reduced their costs by 90% as well as increasing scalability of their monitoring tool.

### Twitter And Wix

Most organisations do massively benefit from a microservice architecture but this also over complicates their system by adding dependencies that are not essential to the functioning of their applications.

Twitter is a massive social media platform with millions of users worldwide which also provides a lot of advertising services as well provides APIs to developers to develop applications making use of it. Twitter's CEO Elon Musk claims that only a mere 20% of microservices in twitter are essential for its functioning while the rest provide little to no functionality or performance improvements.[[8]](#elon-musk-on-twitter)

Wix is another such instance where migrating to microservices proved to be difficult. Wix began its migration to microservices from its monolithic system when they discovered that bugs in one part of the system had the potential to bring down the entire application. They encountered new issues like addressing failures and debugging which were not the case in their monolith system. They had create an entire framework to work with microservices. The employees at Wix also had to face challenges to migrate to microservices from monolith which they were used to which resulted in a culture shift.

## 4.4 Summary Of Challenges

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

In a time characterized by our collective pursuit of convenience and adaptability, the **RentWave** project emerges as a groundbreaking platform, poised to reshape the entire landscape of rental services. RentWave functions as an all-encompassing nexus for a wide spectrum of commodities, offering a comprehensive and versatile solution to individuals in search of renting diverse items. These encompass everything from books and vehicles to even spare guest rooms. The project's core objective revolves around the creation of a resilient, full-stack web application, meticulously structured around the principles of microservices. Each microservice, in this innovative approach, employs a distinct technological stack, mirroring the real-world dynamics of contemporary application development.

The cornerstone of this visionary endeavor is the adoption of an event-driven architecture, which introduces an unparalleled level of dynamism and responsiveness. By combining this architectural framework with the utilization of cutting-edge cloud solutions and the implementation of DevOps best practices, RentWave surges forward as a frontrunner in the domain of modern software design. This amalgamation of advanced technologies and forward-thinking methodologies positions RentWave as an industry disruptor, setting new standards for the seamless and adaptable experience that today's consumers demand. In an age when convenience and adaptability reign supreme, RentWave stands as an embodiment of the innovation that propels us towards a future where renting items of all kinds is easier, more efficient, and more accessible than ever before.

## 5.1 Key Challenges Addressed

### Data Ownership In Microservices

Firstly, dealing with the complexity of building the platform is a big challenge. Making a platform that can handle many types of rentals and user features can be really hard. To tackle this, RentWave uses a microservices approach. Each microservice focuses on a specific part of the platform, which makes it easier to develop and maintain, making the whole process more straightforward and understandable.

### Event Driven Architecture

In addition, the imperative need for real-time interaction demands a strategically engineered solution. Users now anticipate instantaneous engagement and up-to-the-minute updates. In response, RentWave strategically employs an event-driven architecture. This architectural decision substantially addresses the task of establishing smooth communication among the diverse components of the platform, thereby ensuring users remain seamlessly connected and continuously informed.

Event-driven architecture bestows several technical advantages upon RentWave. It fosters an environment where components of the platform can respond to events, such as user actions or external triggers, in real-time. This dynamic responsiveness enables immediate updates and feedback to users, enhancing their experience. Moreover, it facilitates the decoupling of various system elements, allowing them to function independently, which not only simplifies development and maintenance but also enhances system resilience.

By utilizing event-driven architecture, RentWave can also scale horizontally to accommodate growing user demands, as each microservice can handle events and workloads independently. This scalability ensures that the platform can seamlessly expand its capacity to meet the demands of a growing user base, without compromising the quality of real-time interactions.

### Development And Deployment Pipeline

The aspect of Deployment and Operations within the RentWave project is a critical element in ensuring the platform's continuous functionality and improvement. To meet this imperative requirement, the RentWave team has harnessed the power of DevOps technologies, specifically leveraging tools such as Docker and GitHub Actions. This strategic adoption of DevOps principles serves not only to streamline the deployment process but also to orchestrate a seamless and efficient operational environment for the platform.

A pivotal facet of this endeavor involves the establishment of a robust Continuous Integration and Continuous Deployment (CI/CD) pipeline, which is indispensable in the modern software development landscape. This CI/CD pipeline plays a multifaceted role in RentWave's development ecosystem. It automates the integration of code changes, ensuring that new features, bug fixes, and enhancements are smoothly incorporated into the platform without disruptions or errors.

# 6. Design Review

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

## 6.2 Overview Of Microservices And Events

RentWave makes use of several microservices in order to fulfill its requirements of authentication, notifications, scalability, etc.

The brief overview of the core microservices used is shown below:

![Microservices Overview](./images/MicroServices%20Architecture.png)

The functionalities offered by each of the above mentioned services will be described in detail in the following section.

## 6.3 Detailed Design

# 7. RentWave Implementation

# 8. Conclusion

# 9. Appendix

# 10. References

1. ##### [What is Microservices Architecture?- Google Cloud](https://cloud.google.com/learn/what-is-microservices-architecture)
2. ##### [Microservice Scalability Challenges and How to Overcome Them - developer.com](https://www.developer.com/design/microservices-scalability-challenges/)
3. ##### [Why and How Netflix, Amazon, and Uber Migrated to Microservices: Learn from Their Experience](https://www.hys-enterprise.com/blog/why-and-how-netflix-amazon-and-uber-migrated-to-microservices-learn-from-their-experience/)
4. ##### [What Led Amazon to its Own Microservices Architecture](https://thenewstack.io/led-amazon-microservices-architecture/)
5. ##### [Completing the Netflix Cloud Migration](https://about.netflix.com/en/news/completing-the-netflix-cloud-migration)
6. ##### [Optimizing the Netflix API](https://netflixtechblog.com/optimizing-the-netflix-api-5c9ac715cf19)
7. ##### [Scaling up the Prime Video audio/video monitoring service and reducing costs by 90%](https://www.primevideotech.com/video-streaming/scaling-up-the-prime-video-audio-video-monitoring-service-and-reducing-costs-by-90)
8. ##### [Elon Musk on Twitter](https://twitter.com/elonmusk/status/1592177471654604800)
9. ##### [Wix.com Migration to Microservices](https://techbeacon.com/app-dev-testing/monolith-microservices-horror-stories-best-practices)