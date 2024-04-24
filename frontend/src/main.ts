import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

createApp(App).mount('#app')
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}