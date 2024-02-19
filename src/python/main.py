import customtkinter as ctk

ctk.set_appearance_mode('system')
ctk.set_default_color_theme('dark-blue')

root = ctk.CTk()
root.geometry("500x400")
root.title("EMatDystenBot")
root.resizable(False,False)

# GRID

root.columnconfigure(0, weight=1)
root.columnconfigure(1, weight=100)
root.rowconfigure(0, weight=1)
root.rowconfigure(1, weight=1)
root.rowconfigure(2, weight=1)

# DELAY

delayFrame = ctk.CTkFrame(master=root)

delayLabel = ctk.CTkLabel(master=delayFrame,text="Loading...")

defaultDelay = 250

def UpdateDelayLabel(delay):
	delayLabel.configure(text="Delay: "+ str(round(delay)) + "ms")

delaySlider = ctk.CTkSlider(master=delayFrame,from_=0,to=1000,number_of_steps=20,command=UpdateDelayLabel)
delaySlider.set(defaultDelay)
UpdateDelayLabel(defaultDelay)

# LIMIT

limitFrame = ctk.CTkFrame(master=root)

# REALISTIC

realismFrame = ctk.CTkFrame(master=root)

# CONTROL PANEL

cpFrame = ctk.CTkFrame(master=root)

cpTitle = ctk.CTkLabel(master=cpFrame,text="Control Panel",font=ctk.CTkFont(size=25))

# CONTROL PANEL TOGGLE
cpToggleFrame = ctk.CTkFrame(master=cpFrame, height=140, fg_color="transparent")

cpToggleDelay = ctk.CTkSwitch(master=cpToggleFrame,text="Delay")
cpToggleDelay.select()
cpToggleLimit = ctk.CTkSwitch(master=cpToggleFrame,text="Limit")
cpToggleLimit.select()
cpToggleRealism = ctk.CTkSwitch(master=cpToggleFrame,text="Realism")
cpToggleRealism.select()

# CONTROL PANEL RUN
cpRunFrame = ctk.CTkFrame(master=cpFrame, height=140)

# PLACING
gridPadX = 10
gridPadY = 10

delayFrame.grid(row = 0, column = 0, sticky = "nsew", padx = gridPadX, pady = gridPadY)
limitFrame.grid(row = 1, column = 0, sticky = "nsew", padx = gridPadX, pady = gridPadY)
realismFrame.grid(row = 2, column = 0, sticky = "nsew", padx = gridPadX, pady = gridPadY)
cpFrame.grid(row = 0, column = 1, sticky = "nsew", padx = gridPadX, pady = gridPadY, rowspan= 3)

delayLabel.place(relwidth=.75,relheight=.25,relx=.125,rely=.25)
delaySlider.place(relwidth=.75,relheight=.175,relx=.125,rely=.5125)
cpTitle.pack(anchor = "n")

cpPadY = 15

cpToggleFrame.pack(anchor = "n", pady = cpPadY, fill = "x")
cpToggleFrame.columnconfigure(0, weight=1)
cpToggleFrame.rowconfigure(0, weight=1)
cpToggleFrame.rowconfigure(1, weight=1)
cpToggleFrame.rowconfigure(2, weight=1)

cpTogglePadY = 2

cpToggleDelay.grid(row = 0, column = 0, pady = cpTogglePadY)
cpToggleLimit.grid(row = 1, column = 0, pady = cpTogglePadY)
cpToggleRealism.grid(row = 2, column = 0, pady = cpTogglePadY)

cpRunFrame.pack(anchor = "n", pady = cpPadY, fill = "x")

root.mainloop()