Artistic project

This project is a study project to write a web application for art history.
It uses Go as backend. This is as pure Go as possible. The mgo MongoDB is used and some Gorilla modules. Project 
is now developed on Go 1.6 (using Win10) and this is the only version tested. I think it should work on at least
version 1.5, maybe even on older one.
On frontend, this is Bootstrap (currently v3) app. It uses jQuery (it is needed by Bootstrap anyway) and some other
jQuery plugins (no need to reinvent the wheel...) are used also: 
- jQuery dataTables for tables with pagination and searching built-in.
- jQuery validate for form validation.

TODO list
 1. Use jQuery dataTables to make tables fantastic.
 2. Use jQuery Validate for form validation.
 3. Revise the Mongo indexes and put them into app initialization procedure.
 4. Implement more complex filters.
 5. Use HTML5 web storage for some of the functionalities (remember pagination page after refresh, filters...). 
 6. Think how to use web sockets to get data on-the-fly (instead of AJAX).
 7. Test & debug the authorization (and password change) procedure; I'm not sure this is 100% right now...
99. Port to Bootstrap 4 (when available).  
