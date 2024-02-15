<template>
    <div class="container mt-5 text-center">
        <h2 class="display-4 mb-4">User Search</h2>
        <form @submit.prevent="searchUsers" class="mb-4">
            <div class="form-group d-flex justify-content-center align-items-center">
                <label for="searchQuery" class="mr-3" style="font-size: 30px; margin: 20px;">Username: </label>
                <input type="text" id="searchQuery" v-model="searchQuery" class="form-control"
                    placeholder="Enter username" :class="{ 'is-invalid': !isUsernameValid() }"/>
                <button type="submit" class="btn btn-sm btn-outline-secondary ml-2" @click="searchUsers"
                    style="margin: 20px; font-size: 30px;"
                    :disabled="!searchQuery || !isUsernameValid()">Search</button>
            </div>
        </form>
        <p v-if="searchExecuted" class="mt-3" style="font-size: 25px;">
            {{ Text }}
        <ul class="list-group list-group-flush">
            <li v-for="user in UserList" :key="user.userId" class="list-group-item">
                <div class="container">
                    {{ user.username }}
                    <button type="button" class="btn btn-secondary"
                        @click="$router.push(`/users/${user.userId}`)">Profile</button>
                </div>


            </li>
        </ul>
        </p>
    </div>
</template>
  
  
<script>
const token = sessionStorage.getItem('authToken');

export default {
    data() {
        return {
            searchQuery: '',
            searchExecuted: false,
            Text: '',
            UserList: [],
        };
    },
    methods: {
        async searchUsers() {
            try {
                const response = await this.$axios.get(`/users/`, {
                    params: { username: this.searchQuery },
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Accept': 'application/json',
                    },
                });
                this.searchExecuted = true;
                this.UserList = response.data.users;
                this.Text = this.UserList === null ? "No user found with that username." : "";
            }
            catch (error) {
                console.error(error, "Error in user search")
                this.searchExecuted = true;
                if (error.response) {
                    const statusCode = error.response.status;
                    this.notBanned = false;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            this.Text = "You have to log in first";
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            this.Text = "No users with such username";
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                            this.Text = "No users with such username";
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },
        isUsernameValid() {

			const usernameRegex = /^[a-zA-Z][\.]{0,1}([\w][\.]{0,1})*$/
			return usernameRegex.test(this.searchQuery) && this.searchQuery.length >= 3 && this.searchQuery.length <= 25
		},
    },
};
</script>

<style scoped></style>
  