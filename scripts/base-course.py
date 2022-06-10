import sys
import json
import re


"""
    Note 1: this is to be used on the output of pre-process.py script
    Note 2: this will only produce a partial list of required courses for each semester
    USE: python base-course.py <jsonfile> > output.json
"""

def reformat(entry):        
    data = {
        'courseNumber': entry["courseNumber"],
        'subject': entry["subject"],
        'sequenceNumber': entry["sequenceNumber"],
        'courseTitle': entry["courseTitle"],
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
            # spring courses
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
            
            # spring courses
            if re.search("^.*01$", entry["term"]):      
                if entry['subject'] == "CSC":
                    if entry['courseNumber'] in ['111', '115', '225', '226', '230', "320", "360", "370", '361', '460']:
                        springCourseList.append(reformat(entry))
                elif entry['subject'] == "SENG":
                    if entry['courseNumber'] in ["265", "275", "310", "321", "371", "401", "411", "468", "474"]:
                        springCourseList.append(reformat(entry))
                elif entry['subject'] == "ECE":
                    if entry['courseNumber'] in ["300", "310", "320", "330", "340", "360", "455", "458"]:
                        springCourseList.append(reformat(entry))
                    
            # summer courses
            if re.search("^.*05$", entry["term"]):
                if entry['subject'] == "CSC":
                    if entry['courseNumber'] in ['115', '225', '226', '230', "320", "360", "370"]:
                        summerCourseList.append(reformat(entry))
                elif entry['subject'] == "SENG":
                    if entry['courseNumber'] in ["265", "275", "310", "426", "440", "499", "275", "310"]:
                        summerCourseList.append(reformat(entry))
                elif entry['subject'] == "ECE":
                    if entry['courseNumber'] in ["216", "220", "250", "260", "299", "310", "499"]:
                        summerCourseList.append(reformat(entry))

    schedule = {
        'fallTermCourses': fallCourseList,
        'springTermCourses': springCourseList,
        'summerTermCourses': summerCourseList
    }
    print(json.dumps(schedule, indent=2))

if __name__ == '__main__':
    main()
