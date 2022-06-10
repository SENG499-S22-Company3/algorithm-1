import sys
import json
import re


"""
    USE: python pre-process.py <jsonfile> > output.json
"""

def reformat(entry):
    for x in entry["meetingsFaculty"]:
        assign = {
            'startDate': x["meetingTime"]["startDate"],
            'endDate': x["meetingTime"]["endDate"],
            'beginTime': x["meetingTime"]["beginTime"],
            'endTime': x["meetingTime"]["endTime"],
            'hoursWeek': x["meetingTime"]["hoursWeek"],
            'sunday': x["meetingTime"]["sunday"],
            'monday': x["meetingTime"]["monday"],
            'tuesday': x["meetingTime"]["tuesday"],
            'wednesday': x["meetingTime"]["wednesday"],
            'thursday': x["meetingTime"]["thursday"],
            'friday': x["meetingTime"]["friday"],
            'saturday': x["meetingTime"]["saturday"]
        }
        
    data = {
        'courseNumber': entry["courseNumber"],
        'subject': entry["subject"],
        'sequenceNumber': entry["sequenceNumber"],
        'courseTitle': entry["courseTitle"],
        'meetingTime': assign
    }

    return data

def main():
    with open(sys.argv[1],encoding='utf8') as f:
        data = json.load(f)

    fallCourseList = []
    summerCourseList = []
    springCourseList = []
    for entry in data:
        if entry != None:
            if entry['scheduleTypeDescription'] != "Lab" and entry['scheduleTypeDescription'] != "Tutorial":
                # fall courses 
                if re.search("^.*09$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['111', '115', '225', '226', '230', '320', '355', '360', '361', '370', '355']:
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", '310', '321', "350", "360", '421', '474']:
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["220", "241", "250", "255", "260", "355", "360", "365", "370", "380", "399"]:
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "MATH":
                        if entry['courseNumber'] in ["100", "109", "110", "122"]:
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "ENGR":
                        if entry['courseNumber'] in ["110", "130"]:
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "CHEM":
                        if entry['courseNumber'] == "101":
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "PHYS":
                        if entry['courseNumber'] == "110":
                            fallCourseList.append(reformat(entry))
                    elif entry['subject'] == "STAT":
                        if entry['courseNumber'] == "260":
                            fallCourseList.append(reformat(entry))
                
                # spring courses
                elif re.search("^.*01$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['111', '115', '225', '226', '230', "320", "360", "370", '361', '460']:
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", "275", "310", "321", "371", "401", "411", "468", "474"]:
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["300", "310", "320", "330", "340", "360", "455", "458"]:
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "MATH":
                        if entry['courseNumber'] in ["101"]:
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "ENGR":
                        if entry['courseNumber'] in ["120", "141"]:
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "CHEM":
                        if entry['courseNumber'] == "150":
                            springCourseList.append(reformat(entry))
                    elif entry['subject'] == "PHYS":
                        if entry['courseNumber'] == "111":
                            springCourseList.append(reformat(entry))
                
                # summer courses
                elif re.search("^.*05$", entry["term"]): 
                    if entry['subject'] == "CSC":
                        if entry['courseNumber'] in ['115', '225', '226', '230', "320", "360", "370"]:
                            summerCourseList.append(reformat(entry))
                    elif entry['subject'] == "SENG":
                        if entry['courseNumber'] in ["265", "275", "310", "426", "440", "499", "275", "310"]:
                            summerCourseList.append(reformat(entry))
                    elif entry['subject'] == "ECE":
                        if entry['courseNumber'] in ["216", "220", "250", "260", "299", "310", "499"]:
                            summerCourseList.append(reformat(entry))
                    elif entry['subject'] == "ECON":
                        if entry['courseNumber'] == "180":
                            summerCourseList.append(reformat(entry))

    schedule = {
        'fallTermCourses': fallCourseList,
        'springTermCourses': springCourseList,
        'summerTermCourses': summerCourseList
    }
    print(json.dumps(schedule, indent=2))

if __name__ == '__main__':
    main()
