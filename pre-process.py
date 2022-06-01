import sys
import json

"""

USE: python pre-process.py <jsonfile> > output.json

"""

def del_fields(entry):
    del entry["id"]
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
    for x in entry["faculty"]:
        del x["bannerId"]
        del x["category"]
        del x["class"]
        del x["courseReferenceNumber"]
        del x["primaryIndicator"]
        del x["term"]
        del x["emailAddress"]

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
                if entry['subject'] == "CSC":
                    if entry['courseNumber'] in ['111', '115', '230', "225", '361', '320', '360', '370', '226', '460']:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "SENG":
                    if entry['courseNumber'] in ["265", "275", "310", "321", "371", "350", "360", "426", "440", "499", "401"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "ECE":
                    if entry['courseNumber'] in ["255", "260", "310", "458", "355", "360", "360", "455"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "MATH":
                    if entry['courseNumber'] in ["100", "109", "110", "101", "122", "360"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "ENGR":
                    if entry['courseNumber'] in ["120", "110", "130", "141"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "CHEM":
                    if entry['courseNumber'] in ["101", "150"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")
                elif entry['subject'] == "PHYS":
                    if entry['courseNumber'] in ["110", "111"]:
                        print(json.dumps(del_fields(entry),indent=2))
                        print(",")

    print("\t null\n]")

if __name__ == '__main__':
    main()