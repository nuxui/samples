package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: gesture,  

        // application identifier name
        name: "org.nuxui.samples.gesture",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        title: "gesture",
        content: {
            type: main.Home,
        },
    },
}
`