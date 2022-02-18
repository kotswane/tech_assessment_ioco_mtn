# tech_assessment_ioco_mtn

Deployment Guide

1) Clone the repo as below with the terminal

   git clone https://github.com/kotswane/tech_assessment_ioco_mtn.git

2) Change directory to tech_assessment_ioco_mtn

3) Use the same terminal to export the following environment variables
	
	export DATABASE_HOST=localhost
	
	export DATABASE_PORT=3306
	
	export DATABASE_USER=root
	
	export DATABASE_PASS=
	
	export DATABASE_NAME=ioco
	
	export ROBOTS_ENDPOINT=https://robotstakeover20210903110417.azurewebsites.net/robotcpu

3) Insure that you have mysql installed and running on your machine and import the following sql file and allow root user to login without password

   ioco.sql

4) run main.go
==============================================================================================

Test Instructions
