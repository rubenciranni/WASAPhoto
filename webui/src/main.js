import { createApp, reactive } from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import PostItem from './components/PostItem.vue'
import PostList from './components/PostList.vue'
import CommentItem from './components/CommentItem.vue'
import UserItem from './components/UserItem.vue'
import UserListModal from './components/UserListModal.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios
app.config.globalProperties.$paginationLimit = 20
app.component("ErrorMsg", ErrorMsg)
app.component("LoadingSpinner", LoadingSpinner)
app.component("PostItem", PostItem)
app.component("PostList", PostList)
app.component("CommentItem", CommentItem)
app.component("UserItem", UserItem)
app.component("UserListModal", UserListModal)

router.beforeEach(async (to, from) => {
    const authToken = axios.defaults.headers.common['Authorization']
    const userId = localStorage.getItem("userId")
    if (!authToken && !userId && to.name !== 'Login') {
        return { name: 'Login' }
    } else if (authToken && to.name === 'Login') {
        return { name: 'Home' }
    } else if (userId && !authToken) {
        axios.defaults.headers.common['Authorization'] = `Bearer ${userId}`
    }
})
app.use(router)
app.mount('#app')
