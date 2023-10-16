# 1. Abstract

- What Are MicroServices and Event Driven Architecture

- Cons Of MicroServices

- Solutions

- Project Findings

- Future Work

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
6. [Design Review](#6-design-review)
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

- **Data Management**: Choose appropriate data storage solutions, such as databases or data stores, that align with the needs of each microservice. Consider polyglot persistence^2 for flexibility.

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

## 4.2 Successful Implementations

## 4.3 Failed Implementations

## 4.4 Summary Of Challenges

# 5. RentWave Overview

# 6. Design Review

## 6.1 Requirements

## 6.2 Overview Of Microservices And Events

## 6.3 Detailed Design

# 7. RentWave Implementation

# 8. Conclusion

# 9. Appendix

# 10. References
