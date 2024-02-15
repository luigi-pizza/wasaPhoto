import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import LogoutView from '../views/LogoutView.vue'
import ProfileView from '../views/ProfileView.vue'
import PostPhotoView from '../views/PostPhotoView.vue'
import UserSearchView from '../views/UserSearchView.vue'
import StreamView from '../views/StreamView.vue'
import ChangeUsernameView from '../views/ChangeUsernameView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView, redirect:'/login'},
		{path: '/login', component: LoginView},
		{path: '/logout', component: LogoutView},
		{path: '/settings', component: ChangeUsernameView},
		{path: '/photos/', component: PostPhotoView},
		{path: '/users/', component: UserSearchView},
		{path: '/users/:userId', component: ProfileView},
		{path: '/users/:userId/stream/', component: StreamView},

	]
})

export default router
