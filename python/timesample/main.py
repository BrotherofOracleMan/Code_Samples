from datetime import *
from dateutil.parser import *
import logging


logging.basicConfig(level=logging.NOTSET)
logger = logging.getLogger()

def CheckDate(Date):
    #Check if Date is valid
    #Check if Date is greater than today 
    datetime_object = None
    if Date is not None:
      logger.info("Validating date")

    try:
        datetime_object = isoparse(Date)
    except ValueError:
        logger.error("Value Error is called Date is invalid")
        return False,"Date is Invalid. Please enter a Valid Date" 
    
    if datetime_object.date() < datetime.today().date():
        logger.error("Date is in the past")
        return False ,"Date should be before. Re enter a event date that is Today or after "
    
    return True,"Valid Date Given"


def CheckTime(Date,time):
    #Check if time is valid
    #Check if Date is valid
    #if both are valid compare it todays time and
    datetime_object = None
    try:
        datetime_object = parse(Date + " " + time)
        logger.info(datetime_object)
    except ValueError:
        logger.error("Invalid time was given. Re Enter a valid time")
        return False, "Re-enter Valid time"

    if datetime_object < datetime.today():
        logger.error("Past time given")
        return False, "Re-enter a time that is in the future"
    return True, "Valid Time Given"
    
def CheckEvent(Event):
    logger.info("Validating Event field")
    if Event == "":
        return False,"Event cannot be empty"
    return True, ""



def main():

    CheckDate_valid , err= CheckDate('2021-08-02')
    CheckTime_valid, err = CheckTime('2021-08-04',"22:00")
    logger.info("Main Loop")


main()