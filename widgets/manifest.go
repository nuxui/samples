package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: widgets,  

        // application identifier name
        name: "org.nuxui.samples.widgets",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width: 50%,
        height: 50%,
        title: "widgets",
        content: {type: main.Home},
    },
}
`