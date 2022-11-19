package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: hello,  

        // application identifier name
        name: "org.nuxui.samples.hello",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width: 15%,
        height: 2:1,
        title: "hello",
        content: {
            type: nuxui.org/nuxui/ui.Text,
			text: "Hello nuxui",
        },
        background: #ffffff,
    },
}
`