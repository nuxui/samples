package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: Gallery,  

        // application identifier name
        name: "org.nuxui.samples.widgets.gallery",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width:  "1:1",
        height: "500px",
        title:  "Gallery",
        content: {
            type: main.Home,
        },
    },
}
`