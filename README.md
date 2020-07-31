# Assignment 5 (SEG 2105 Z)

Extension/Revamp of the Pandemic System (Assignment #2)  

**Description:** a simple system that helps test providers to administer viral tests and for lab technicians to record test results for patients.  


## Design & Architecture

### Design Process
A mixture of top-down and bottom-up design, but with a more central focus on bottom-up design.
I knew I wanted the overall structure to be MVC, so I built the small individual components first, and I would start arranging, organizing, and adjusting sections of the system to be fit into an MVC architecture.



### 1st Iteration: build a minimum viable system based on the following user stories
#### User Stories
1. User enters personal information on patient information page
2. Test Provider can view available users on a patient repository and select a user for a test
3. Lab Technician accesses patient repository to record test results
4. User can login and see their test results


### Tech Stack
A list of all the programming languages, technology, and frameworks that I used:  
* Go 1.14 (no front/back-end web framework because go comes with built-in a web server and web templating library)
* CSS
* HTML
* Bootstrap (CSS Framework)

I chose Go as the main language for developing this project because I had worked with it in my CSI2120 class and I found it interesting, so I wanted to get better at using it.