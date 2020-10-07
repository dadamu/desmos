(window.webpackJsonp=window.webpackJsonp||[]).push([[45],{392:function(t,e,n){"use strict";n.r(e);var s=n(42),a=Object(s.a)({},(function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[n("h1",{attrs:{id:"installing-and-running-a-desmos-fullnode"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#installing-and-running-a-desmos-fullnode"}},[t._v("#")]),t._v(" Installing and running a Desmos fullnode")]),t._v(" "),n("div",{staticClass:"custom-block warning"},[n("p",{staticClass:"custom-block-title"},[t._v("This guide is for new fullnodes only")]),t._v(" "),n("p",[t._v("If you have previously run a fullnode and you wish to update it instead, please follow the "),n("RouterLink",{attrs:{to:"/fullnode/update.html"}},[t._v("updating guide")]),t._v(".")],1)]),t._v(" "),n("h2",{attrs:{id:"requirements"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#requirements"}},[t._v("#")]),t._v(" Requirements")]),t._v(" "),n("h3",{attrs:{id:"understanding-pruning"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#understanding-pruning"}},[t._v("#")]),t._v(" Understanding pruning")]),t._v(" "),n("p",[t._v("In order to run a full node, different hardware requirements should be met based on the pruning strategy you would like to use.")]),t._v(" "),n("p",[n("em",[t._v("Pruning")]),t._v(" is the term used to identify the periodic action that can be taken in order to free some disk space on your full node. This is done by removing old blocks data from the disk, freeing up space.")]),t._v(" "),n("p",[t._v("Inside Desmos, there are various types of pruning strategies that can be applied. The main ones are:")]),t._v(" "),n("ul",[n("li",[n("p",[n("code",[t._v("default")]),t._v(": the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals")])]),t._v(" "),n("li",[n("p",[n("code",[t._v("nothing")]),t._v(": all historic states will be saved, nothing will be deleted (i.e. archiving node)")])]),t._v(" "),n("li",[n("p",[n("code",[t._v("everything")]),t._v(": all saved states will be deleted, storing only the current state; pruning at 10 block intervals")])])]),t._v(" "),n("h3",{attrs:{id:"hardware-requirements"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#hardware-requirements"}},[t._v("#")]),t._v(" Hardware requirements")]),t._v(" "),n("p",[t._v("You can easily understand how using a pruning strategy of "),n("code",[t._v("nothing")]),t._v(" will use more disk space than "),n("code",[t._v("everything")]),t._v(". For this reason, there are different disk space that we recommend based on the pruning strategy you choose:")]),t._v(" "),n("table",[n("thead",[n("tr",[n("th",{staticStyle:{"text-align":"center"}},[t._v("Pruning strategy")]),t._v(" "),n("th",{staticStyle:{"text-align":"center"}},[t._v("Minimum disk space")]),t._v(" "),n("th",{staticStyle:{"text-align":"center"}},[t._v("Recommended disk space")])])]),t._v(" "),n("tbody",[n("tr",[n("td",{staticStyle:{"text-align":"center"}},[n("code",[t._v("nothing")])]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("20 GB")]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("40 GB")])]),t._v(" "),n("tr",[n("td",{staticStyle:{"text-align":"center"}},[n("code",[t._v("default")])]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("80 GB")]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("120 GB")])]),t._v(" "),n("tr",[n("td",{staticStyle:{"text-align":"center"}},[n("code",[t._v("everything")])]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("120 GB")]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("> 240 GB")])])])]),t._v(" "),n("p",[t._v("A part from disk space, the following requirements should be met.")]),t._v(" "),n("table",[n("thead",[n("tr",[n("th",{staticStyle:{"text-align":"center"}},[t._v("Minimum CPU cores")]),t._v(" "),n("th",{staticStyle:{"text-align":"center"}},[t._v("Recommended CPU cores")])])]),t._v(" "),n("tbody",[n("tr",[n("td",{staticStyle:{"text-align":"center"}},[t._v("2")]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("4")])])])]),t._v(" "),n("table",[n("thead",[n("tr",[n("th",{staticStyle:{"text-align":"center"}},[t._v("Minimum RAM")]),t._v(" "),n("th",{staticStyle:{"text-align":"center"}},[t._v("Recommended RAM")])])]),t._v(" "),n("tbody",[n("tr",[n("td",{staticStyle:{"text-align":"center"}},[t._v("4 GB")]),t._v(" "),n("td",{staticStyle:{"text-align":"center"}},[t._v("8 GB")])])])]),t._v(" "),n("h2",{attrs:{id:"_1-setup-your-environment"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#_1-setup-your-environment"}},[t._v("#")]),t._v(" 1. Setup your environment")]),t._v(" "),n("p",[t._v("In order to run a fullnode, you need to build "),n("code",[t._v("desmosd")]),t._v(" and "),n("code",[t._v("desmoscli")]),t._v(" which require "),n("code",[t._v("Go")]),t._v(", "),n("code",[t._v("git")]),t._v(", "),n("code",[t._v("gcc")]),t._v(" and "),n("code",[t._v("make")]),t._v(" installed.")]),t._v(" "),n("p",[t._v("This process depends on your working environment.")]),t._v(" "),n("tabs",[n("tab",{attrs:{name:"Linux"}},[n("p",[t._v("The following example is based on "),n("strong",[t._v("Ubuntu (Debian)")]),t._v(" and assumes you are using a terminal environment by default. Please run the equivalent commands if you are running other Linux distributions.")]),t._v(" "),n("div",{staticClass:"language-bash line-numbers-mode"},[n("pre",{pre:!0,attrs:{class:"language-bash"}},[n("code",[n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Install git, gcc and make")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("sudo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("apt")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("install")]),t._v(" build-essential --yes\n\n"),n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Install Go with Snap")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("sudo")]),t._v(" snap "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("install")]),t._v(" go --classic\n\n"),n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Export environment variables")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export GOPATH=\""),n("span",{pre:!0,attrs:{class:"token environment constant"}},[t._v("$HOME")]),t._v("/go\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export GOBIN=\""),n("span",{pre:!0,attrs:{class:"token variable"}},[t._v("$GOPATH")]),t._v("/bin\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export PATH=\""),n("span",{pre:!0,attrs:{class:"token variable"}},[t._v("$GOBIN")]),t._v(":"),n("span",{pre:!0,attrs:{class:"token environment constant"}},[t._v("$PATH")]),t._v("\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("source")]),t._v(" ~/.profile\n")])]),t._v(" "),n("div",{staticClass:"line-numbers-wrapper"},[n("span",{staticClass:"line-number"},[t._v("1")]),n("br"),n("span",{staticClass:"line-number"},[t._v("2")]),n("br"),n("span",{staticClass:"line-number"},[t._v("3")]),n("br"),n("span",{staticClass:"line-number"},[t._v("4")]),n("br"),n("span",{staticClass:"line-number"},[t._v("5")]),n("br"),n("span",{staticClass:"line-number"},[t._v("6")]),n("br"),n("span",{staticClass:"line-number"},[t._v("7")]),n("br"),n("span",{staticClass:"line-number"},[t._v("8")]),n("br"),n("span",{staticClass:"line-number"},[t._v("9")]),n("br"),n("span",{staticClass:"line-number"},[t._v("10")]),n("br"),n("span",{staticClass:"line-number"},[t._v("11")]),n("br")])])]),t._v(" "),n("tab",{attrs:{name:"MacOS"}},[n("p",[t._v("To install the required build tools, simple "),n("a",{attrs:{href:"https://apps.apple.com/hk/app/xcode/id497799835?l=en&mt=12",target:"_blank",rel:"noopener noreferrer"}},[t._v("install Xcode from the Mac App Store"),n("OutboundLink")],1),t._v(".")]),t._v(" "),n("p",[t._v("To install "),n("code",[t._v("Go")]),t._v(" on "),n("strong",[t._v("MacOS")]),t._v(", the best option is to install with "),n("a",{attrs:{href:"https://brew.sh/",target:"_blank",rel:"noopener noreferrer"}},[n("strong",[t._v("Homebrew")]),n("OutboundLink")],1),t._v(". To do so, open the "),n("code",[t._v("Terminal")]),t._v(" application and run the following command:")]),t._v(" "),n("div",{staticClass:"language-bash line-numbers-mode"},[n("pre",{pre:!0,attrs:{class:"language-bash"}},[n("code",[n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Install Homebrew")]),t._v("\n/usr/bin/ruby -e "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v('"'),n("span",{pre:!0,attrs:{class:"token variable"}},[n("span",{pre:!0,attrs:{class:"token variable"}},[t._v("$(")]),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("curl")]),t._v(" -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install"),n("span",{pre:!0,attrs:{class:"token variable"}},[t._v(")")])]),t._v('"')]),t._v("\n")])]),t._v(" "),n("div",{staticClass:"line-numbers-wrapper"},[n("span",{staticClass:"line-number"},[t._v("1")]),n("br"),n("span",{staticClass:"line-number"},[t._v("2")]),n("br")])]),n("blockquote",[n("p",[t._v("If you don't know how to open a "),n("code",[t._v("Terminal")]),t._v(", you can search it by typing "),n("code",[t._v("terminal")]),t._v(" in "),n("code",[t._v("Spotlight")]),t._v(".")])]),t._v(" "),n("p",[t._v("After "),n("strong",[t._v("Homebrew")]),t._v(" is installed, run")]),t._v(" "),n("div",{staticClass:"language-bash line-numbers-mode"},[n("pre",{pre:!0,attrs:{class:"language-bash"}},[n("code",[n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Install Go using Homebrew")]),t._v("\nbrew "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("install")]),t._v(" go\n\n"),n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Install Git using Homebrew")]),t._v("\nbrew "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("install")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token function"}},[t._v("git")]),t._v("\n\n"),n("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Export environment variables")]),t._v("\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export GOPATH=\""),n("span",{pre:!0,attrs:{class:"token environment constant"}},[t._v("$HOME")]),t._v("/go\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export GOBIN=\""),n("span",{pre:!0,attrs:{class:"token variable"}},[t._v("$GOPATH")]),t._v("/bin\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("echo")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token string"}},[t._v("'export PATH=\""),n("span",{pre:!0,attrs:{class:"token variable"}},[t._v("$GOBIN")]),t._v(":"),n("span",{pre:!0,attrs:{class:"token environment constant"}},[t._v("$PATH")]),t._v("\"'")]),t._v(" "),n("span",{pre:!0,attrs:{class:"token operator"}},[t._v(">>")]),t._v(" ~/.profile\n"),n("span",{pre:!0,attrs:{class:"token builtin class-name"}},[t._v("source")]),t._v(" ~/.profile\n")])]),t._v(" "),n("div",{staticClass:"line-numbers-wrapper"},[n("span",{staticClass:"line-number"},[t._v("1")]),n("br"),n("span",{staticClass:"line-number"},[t._v("2")]),n("br"),n("span",{staticClass:"line-number"},[t._v("3")]),n("br"),n("span",{staticClass:"line-number"},[t._v("4")]),n("br"),n("span",{staticClass:"line-number"},[t._v("5")]),n("br"),n("span",{staticClass:"line-number"},[t._v("6")]),n("br"),n("span",{staticClass:"line-number"},[t._v("7")]),n("br"),n("span",{staticClass:"line-number"},[t._v("8")]),n("br"),n("span",{staticClass:"line-number"},[t._v("9")]),n("br"),n("span",{staticClass:"line-number"},[t._v("10")]),n("br"),n("span",{staticClass:"line-number"},[t._v("11")]),n("br")])])]),t._v(" "),n("tab",{attrs:{name:"Windows"}},[n("p",[t._v("The software has not been tested on "),n("strong",[t._v("Windows")]),t._v(". If you are currently running a "),n("strong",[t._v("Windows")]),t._v(" PC, the following options should be considered:")]),t._v(" "),n("ol",[n("li",[t._v("Switch to a "),n("strong",[t._v("Mac")]),t._v(" 👨‍💻.")]),t._v(" "),n("li",[t._v("Wipe your hard drive and install a "),n("strong",[t._v("Linux")]),t._v(" system on your PC.")]),t._v(" "),n("li",[t._v("Install a separate "),n("strong",[t._v("Linux")]),t._v(" system using "),n("a",{attrs:{href:"https://www.virtualbox.org/wiki/Downloads",target:"_blank",rel:"noopener noreferrer"}},[t._v("VirtualBox"),n("OutboundLink")],1)]),t._v(" "),n("li",[t._v("Run a "),n("strong",[t._v("Linux")]),t._v(" instance on a cloud provider.")])]),t._v(" "),n("p",[t._v("Note that is still possible to build and run the software on "),n("strong",[t._v("Windows")]),t._v(" but it may give you unexpected results and it may require additional setup to be done. If you insist to build and run the software on "),n("strong",[t._v("Windows")]),t._v(", the best bet would be installing the "),n("a",{attrs:{href:"https://chocolatey.org/",target:"_blank",rel:"noopener noreferrer"}},[t._v("Chocolatey"),n("OutboundLink")],1),t._v(" package manager.")])])],1),t._v(" "),n("h2",{attrs:{id:"_2-install-the-software"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#_2-install-the-software"}},[t._v("#")]),t._v(" 2. Install the software")]),t._v(" "),n("p",[t._v("Once you have setup your environment correctly, you are now ready to install the Desmos software and start your full node.")]),t._v(" "),n("p",[t._v("In order to do so, you have two options:")]),t._v(" "),n("ol",[n("li",[n("p",[n("RouterLink",{attrs:{to:"/fullnode/setup/automatic.html"}},[t._v("Automatic installation")]),t._v("."),n("br"),t._v("\nThis requires you to run a few commands so that you can easily have a Desmos full node running in a couple of minutes. This is recommended to most users that do not yet have the skills to go through a manual setup.")],1)]),t._v(" "),n("li",[n("p",[n("RouterLink",{attrs:{to:"/fullnode/setup/manual.html"}},[t._v("Manual installation")]),t._v("."),n("br"),t._v("\nThis is recommended to the those who want to understand each step or want to customize their setup accordingly to their needs. It is not recommended to people running a validator for the first time, although everyone should take a look at it once.")],1)])])],1)}),[],!1,null,null,null);e.default=a.exports}}]);