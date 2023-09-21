# OneCV Golang Backend Api Technical Assessment
Name: Muhammad Jabir Bin Jamalul Ashik

## Hosted Repository
URL Link: https://g0hbpexwlj.execute-api.us-east-2.amazonaws.com/

## Assumptions Made
- All user stories endpoints assume that the respective student and teacher records exist in the database

## Implemented API Endpoints
User Stories Endpoints
- GET: /api/commonstudents
- POST: /api/register
- POST: /api/suspend
- POST: /api/retrievefornotifications

In addition to the User stories:
- GET: /api/admin/student?student= - Retrieve the list of students based on query parameter, all students if empty, else, specified students
Response Code: 200
Response body: ```{ "students": [] array of students } ```
- POST: /api/admin/student - Post a new student in the database
Request body: ``` { "email": string } ```
Response Code: 204
Request body: ``` { "student": string } ```
- GET: /api/admin/teacher?teacher= - Retrieve the list of teachers based on query parameter, all teachers if empty, else, specified teachers
Rsponse Code: 200
Response body: ```{ "teachers": [] array of teachers } ```

- POST: /api/admin/teacher - Post a new teacher in the database
Response Code: 204
Request body: ``` { "email": string } ```
Request body: ``` { "teacher": string } ```

## Testing in Local Development Environment
Ensure that your docker daemon is running, keep ports 8080 and 5432 unallocated
```sh
docker compose up
```
The api will be running on http://localhost:8080
