import path
import os
class AuditManager :
    max_entries_per_file:int
    directory_name:str
    def __init__(self,max_entries_per_file:int,directory_name:str):
        self.max_entries_per_file = max_entries_per_file
        self.directory_name = directory_name
        

    def add_record(self,visitor_name:str,time_of_visit):
        file_paths = os.listdir(self.directory_name)
        file_paths.sort()
        new_record = visitor_name + ";" + time_of_visit
        if len(file_paths) == 0:
            new_file = path.join(self.directory_name,"audit_1.txt")
            open(new_file,"w") as f:
                f.write()


am = AuditManager(10,"./")
print(am.add_record("",""))
