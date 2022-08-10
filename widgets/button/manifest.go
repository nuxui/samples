package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: button,  

        // application identifier name
        name: "org.nuxui.samples.widgets.button",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width:  "18%",
        height: "2:1",
        title:  "button",
        content: {
            type: main.Home,
        },
    },
}
`