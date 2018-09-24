package model


var CreateScanQuery = "INSERT INTO scans VALUES (NULL,?,?,?,?)"
var GetIdByScanNameQuery = "SELECT scan_id FROM scans WHERE scan_name = ?"
var InsertPersonQuery = "INSERT INTO persons VALUES (NULL,?,?,?,?)"
var GetPersonsFromScanNameQuery = "SELECT full_name from persons join scans on persons.scan_id = scans.scan_id	where scan_name = ?;"
var GetALLScans = "SELECT * from scans;"
var SelectScanNames = "SELECT scan_name from scans;"