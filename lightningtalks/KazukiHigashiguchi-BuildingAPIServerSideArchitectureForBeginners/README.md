# Building API Server-side Architecture for beginners

Kazuki Higashiguchi ([@hgsgtk](https://github.com/hgsgtk))

* [Slides](slides.pdf)

## Abstract

When we creates an application, we often need to consider about layer design. However, it’s difficult for beginners to design the architecture from beginning. To overcome difficulties, I will talk about the method of gradually growing our architecture with the growth of the members.

## Detail
### Problem of considering architecture for beginners 

In considering the server side architecture, the judge is greatly influenced by the situation such as the request for the application and team background.
I did the development under the following circumstances

- FinTech startup
- Launch a new service
- No member has experience of using Go language in business

Under such circumstances, we must get feedback to design from the practical results in a quick cycle. 
Also, we must understand the Go language knowledge like context handling.

### Strategy to solve the problem 

It is often that important factors are not available, that is development know-how, developer skill, all requirements of service.
It is necessary to set the image to be aimed for us, and explore how to implement it.
For the purpose, we proceeded with making it a common recognition that “building and breaking”.

In addition, we put emphasis on unit testing to keep getting feedback on the way that our application should be in a quick cycle by practicing.  Because we aim to get notices about design (such as degree of coupling) by unit testing.

### gradually growing architecture 
I will talk about the process of gradually growing the architecture. It’s the process of practicing the following steps in phases.

- get used to API development in Go language
- how to use interface to write Go-like code
- write good testable code
- represent complex business logic in code

Also, I will mention the following points in the process.

- CI environment to prepare at first
- Anti-pattern which tends to stepped from beginner to intermediate level
- maintenance of test helpers not to compromise development experience
