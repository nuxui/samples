package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: formater,  

        // application identifier name
        name: "org.nuxui.samples.formater",
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
        title:  "Image Formater",
        content: {
            type: main.Home,
        },
        background: #ffffff,
    },
}
`