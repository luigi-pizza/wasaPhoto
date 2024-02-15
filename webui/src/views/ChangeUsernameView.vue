<template>
    <div class="changeUsername-container">
        <form @submit.prevent="submitForm" class="changeUsername-form">
            <h2 class="mb-4">Change Username</h2>
            <div class="mb-3">
                <label for="inputName" class="form-label">New Name</label>
                <input type="text" class="form-control" id="newname" aria-describedby="usernameHelp"
				    v-model="newname" :class="{ 'is-invalid': !isUsernameValid() }">
            </div>
            <button type="submit" class="btn btn-primary" :disabled="!newname || !isUsernameValid()">Submit</button>
            <div class="alert alert-success" role="alert" v-if="changedSuccess" style="margin: 10px;">
                Name changed successfully!
            </div>
            <ErrorMsg :msg="error_msg" v-else-if="errore" style="margin: 10px;"/>
        </form>
    </div>
</template>
  
<script>
import ErrorMsg from '@/components/ErrorMsg.vue'
const token = sessionStorage.getItem('authToken');
export default {
    components: {
        ErrorMsg
    },
    data() {
        return {
            newname: '',
            changedSuccess: false,
            errore: false,
            error_msg: '',
        };
    },
    methods: {
        async submitForm() {
            try {
                const config = {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                };
                const response = await this.$axios.put(`/settings/username`, { username: this.newname }, config);
                this.changedSuccess = true;
                this.errore = false;
            }
            catch (error) {
                console.error(error, "Error in changing name");
                const statusCode = error.response.status;
                switch (statusCode) {
                    case 401:
                        console.error('Access Unauthorized:', error.response.data);
                        this.error_msg = "You are not logged in"
                        break;
                    case 400:
                        console.error('Bad request:', error.response.data);
                        this.error_msg = "Name already in use"
                        break;
                    case 403:
                        console.error('Forbidden Action: ', error.response.data);
                        this.error_msg = "Username already in use"
                        break
                    case 404:
                        console.error('Not found: ', error.response.data);
                        this.error_msg = "You are not logged in"
                    default:
                        console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                        this.error_msg = "You should login first"
                }
                this.changedSuccess = false;
                this.errore = true;
            }

            this.newname = '';

        },
        isUsernameValid() {

			const usernameRegex = /^[a-zA-Z][\.]{0,1}([\w][\.]{0,1})*[\w]$/
			return usernameRegex.test(this.newname) && this.newname.length >= 5 && this.newname.length <= 25
		},
    },
};
</script>

<style scoped>
.changeUsername-container {
    display: flex;
    justify-content: center;
    text-align:center;
    align-items: center;
    height: 100vh;
}

.changeUsername-form {
    padding: 20px;
    border: none;
    align-items: center;
    border-radius: 8px;
}

.changeUsername-label {
    padding: 3px;
    display: block;
    margin-bottom: 8px;
}

.changeUsername-btn {
    align-self: center;
}

</style>
  