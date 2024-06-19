#!/bin/bash

# To run the script as postgres user every hour, use the following command:
# sudo -u postgres crontab -e
# 0 * * * * /bin/bash /home/postgres/pg_backup.sh


backup_dir="/home/postgres/backups"
backup_filename="$(date +'%Y%m%d_%H%M%S').sql"

if [ ! -d "$backup_dir" ]; then
    mkdir -p "$backup_dir"
fi

/usr/bin/pg_dump -d banague -F plain -f "$backup_dir/$backup_filename"

if [ $? -eq 0 ]; then
    echo "Backup created successfully: $backup_filename" | tee -a /home/postgres/backups/backup.log
else
    echo "Backup failed" | tee -a /home/postgres/backups/backup.log
fi