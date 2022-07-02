import sys
import json
import re


"""
    USE: python pre-process.py <jsonfile> > output.json
"""

def main():
    with open(sys.argv[1],encoding='utf8') as f:
        data = json.load(f)

    fallTerm = []
    springTerm = []
    summerTerm = []

    for entry in data['hardScheduled']['fallTermCourses']:
        if entry['subject'] in ["MATH", "ENGR", "CHEM", "PHYS", "ECON"]:
            fallTerm.append(entry)

    for entry in data['hardScheduled']['springTermCourses']:
        if entry['subject'] in ["MATH", "ENGR", "CHEM", "PHYS", "ECON"]:
            springTerm.append(entry)

    for entry in data['hardScheduled']['summerTermCourses']:
        if entry['subject'] in ["MATH", "ENGR", "CHEM", "PHYS", "ECON"]:
            summerTerm.append(entry)

    input = {
        "hardScheduled": {
            "fallTermCourses": fallTerm,
            "springTermCourses": springTerm,
            "summerTermCourses": summerTerm
        },
        "coursesToSchedule": data['coursesToSchedule'],
        "professors": data['professors']
    }



    print(json.dumps(input, indent=2))
                


if __name__ == '__main__':
    main()
