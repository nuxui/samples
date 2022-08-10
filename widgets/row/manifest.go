package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: row,  

        // application identifier name
        name: "org.nuxui.samples.widgets.row",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width:  "200px",
        height: "200px",
        title:  "row",
        content: {
            type: main.Home,
        },
    },
}
`