<template>
    <div class="login-container">
        <LoadingSpinner v-if="loading"></LoadingSpinner>
        <div class="login-form">
            <h2>Login</h2>
            <form @submit.prevent="login">
                <label class="login-label" for="username">Username:</label>
                <input type="text" class="form-control" id="username" aria-describedby="usernameHelp" 
							v-model="username" :class="{ 'is-invalid': !isUsernameValid() }">
                <button type="submit" class="btn btn-sm btn-outline-primary login-btn" :disabled="!username || !isUsernameValid() || loading"
                    style="align-self: center; float: center; font-size: 20px;">Login <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#key" />
                    </svg></button>
            </form>
            <div v-if="identifier !== null">
                <p>Login successful! User identifier: {{ identifier.userId }}</p>
            </div>
        </div>
    </div>
</template>
  
<script>
const token = sessionStorage.getItem('authToken');
export default {
    data() {
        return {
            username: "",
            identifier: null,
            loading: false,
        };
    },

    methods: {
        async login() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post('/login', { username: this.username }, {
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                    },
                });
                console.log(response)
                this.identifier = response.data
                this.SaveToSessionStorage()
            } catch (error) {
                console.error("Error while logging in!");
            }
            this.loading = false;
            this.navigateToMyPage()
        },
        navigateToMyPage() {
            localStorage.setItem('loginExecuted', '1')
            location.reload();
            // this.$router.push('/users/' + this.identifier.userId);
        },
        SaveToSessionStorage() {
            const bearerToken = `${this.identifier.userId}`;
            sessionStorage.setItem('authToken', bearerToken);
            sessionStorage.setItem('username', this.identifier.username);
            sessionStorage.setItem('userId', this.identifier.userId);
        },
        isUsernameValid() {

			const usernameRegex = /^[a-zA-Z][\.]{0,1}([\w][\.]{0,1})*[\w]$/
			return usernameRegex.test(this.username) && this.username.length >= 5 && this.username.length <= 25
		},
    },
};
</script>
  
<style scoped>
.login-container {
    display: flex;
    justify-content: center;
    text-align:center;
    align-items: center;
    height: 100vh;
}

.login-form {
    padding: 20px;
    border: none;
    align-items: center;
    border-radius: 8px;
}

.login-label {
    padding: 3px;
    display: block;
    margin-bottom: 8px;
}

.login-btn {
    align-self: center;
}

</style>
  