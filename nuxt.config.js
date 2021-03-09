export default {
    ssr: false,
    target: 'static',
    head: {
        title: 'Recipes',
        htmlAttrs: {
            lang: 'en'
        },
        meta: [
            { charset: 'utf-8' },
            {
                name: 'viewport',
                content: 'width=device-width, initial-scale=1'
            },
            { hid: 'description', name: 'description', content: '' }
        ],
        link: [
            {
                rel: 'stylesheet',
                type: 'text/css',
                href: 'https://unpkg.com/mvp.css'
            }
        ]
    },
    css: [],
    plugins: [],
    components: false,
    buildModules: [],
    modules: [],
    build: {
        transpile: []
    },
    server: {
        port: 3000,
        timing: false
    },
    buildDir: 'dist'
};