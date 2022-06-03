import sys
import json
import re


"""

USE: python pre-process.py <jsonfile> > output.json

"""

def del_fields(entry):
    del entry["id"]
    del entry["termDesc"]
    del entry["courseReferenceNumber"]
    del entry["partOfTerm"]
    del entry["subjectDescription"]
    del entry["campusDescription" ]
    del entry["scheduleTypeDescription"]
    del entry["creditHours"]
    del entry["maximumEnrollment"]
    del entry["enrollment"]
    del entry["seatsAvailable"]
    del entry["waitCapacity"]
    del entry["waitCount"]
    del entry["waitAvailable"]
    del entry["crossList"]
    del entry["crossListCapacity"]
    del entry["crossListCount"]
    del entry["crossListAvailable"]
    del entry["creditHourHigh"]
    del entry["creditHourLow"]
    del entry["creditHourIndicator"]
    del entry["openSection"]
    del entry["linkIdentifier"]
    del entry["isSectionLinked"]
    del entry["subjectCourse"]
    del entry["faculty"]

    for x in entry["meetingsFaculty"]:
        del x["category"]
        del x["class"]
        del x["courseReferenceNumber"]
        del x["faculty"]
        del x["meetingTime"]["building"]
        del x["meetingTime"]["buildingDescription"]
        del x["meetingTime"]["campus"]
        del x["meetingTime"]["category"]
        del x["meetingTime"]["class"]
        del x["meetingTime"]["courseReferenceNumber"]
        del x["meetingTime"]["meetingScheduleType"]
        del x["meetingTime"]["meetingType"]
        del x["meetingTime"]["room"]
        del x["meetingTime"]["term"]
        del x["meetingTime"]["campusDescription"]
        del x["meetingTime"]["meetingTypeDescription"]
        del x["meetingTime"]["campusDescription"]
        del x["term"]
        
    del entry["reservedSeatSummary"]
    del entry["sectionAttributes"]
    del entry["instructionalMethod"]
    del entry["instructionalMethodDescription"]

    return entry

def main():
    with open(sys.argv[1],encoding='utf8') as f:
        data = json.load(f)

    print("[")
    for entry in data:
        if entry != None:
            
            if entry['scheduleTypeDescription'] != "Lab" and entry['scheduleTypeDescription'] != "Tutorial":
                # fall courses 
                if re.search("^.*09$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['111', '115', '225', '226', '230', '320', '355', '360', '361', '370', '355']:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", '310', '321', "350", "360", '421', '474']:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["255", "260", "355"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "MATH":
                        if entry['courseNumber'] in ["100", "109", "110", "122"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ENGR":
                        if entry['courseNumber'] in ["110", "130"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "CHEM":
                        if entry['courseNumber'] == "101":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "PHYS":
                        if entry['courseNumber'] == "110":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "STAT":
                        if entry['courseNumber'] == "260":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                
                # spring courses
                elif re.search("^.*01$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['111', '115', '225', '226', '230', "320", "360", "370", '361', '460']:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", "275", "310", "321", "371", "401", "411", "468", "474"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["458", "360", "455"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "MATH":
                        if entry['courseNumber'] in ["101"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ENGR":
                        if entry['courseNumber'] in ["120", "141"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "CHEM":
                        if entry['courseNumber'] == "150":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "PHYS":
                        if entry['courseNumber'] == "111":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                
                # summer courses
                elif re.search("^.*05$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['115', '225', '226', '230', "320", "360", "370"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", "275", "310", "426", "440", "499", "275", "310"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["310"]:
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")
                    elif entry['subject'] == "ECON":
                        if entry['courseNumber'] == "180":
                            print(json.dumps(del_fields(entry),indent=2))
                            print(",")

    print("\t null\n]")

if __name__ == '__main__':
    main()