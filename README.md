# A Simple GUI for Rogue SSD's Data Cloaking Demo

This is a simple graphic UI for the RS data cloaking demo. For the demo to
run properly, it is important to set up the initial state of the machine 
as follows.

- The RS SSD device is NOT used for OS boot
- Visible namespace is /dev/nvme0n1 and formatted in ext4
- Visible namespace's mount point is /mnt/nvme, mounted.
- Hidden namespace is /dev/nvme0n2, and formatted in ext4
- hidden namespace's mount point /mnt/nvme2 exists, but NOT mounted

# Install dependencies
On Ubuntu (tested version 16.04), do 
```
sudo apt-get install build-essential libgtk-3-dev
go get github.com/andlabs/ui
```
# Build GUI application
Use `make build` or simple `make` to build the GUI.

# Run GUI application

Run `./openclose` as root.
