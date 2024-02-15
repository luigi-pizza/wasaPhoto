<template>
  <div class="container mt-5">
    <h2 class="display-4 mb-4">Upload Photo</h2>
    <form @submit.prevent="uploadPhoto" class="needs-validation" novalidate style="font-size: 17px;">
      <div class="container mt-3">
        <div class="input-group">
          <input type="file" class="form-control-file" accept=".png" id="photo" @change="onFileChange" required />
          <div class="invalid-feedback" v-if="!photo">Photo is required</div>
        </div>
      </div>


      <div class="row" style="margin: 10px;">
        <label for="caption">Caption:</label>
        <textarea class="form-control" id="caption" v-model="caption"></textarea>
      </div>

      <button type="submit" class="btn btn-primary btn-lg" style="margin: 20px;">Upload</button>
      <p v-if="uploadSuccess" class="alert alert-success mt-3">{{ endText }}</p>
    </form>
  </div>
</template>


<script>
const token = sessionStorage.getItem('authToken');
export default {
  data() {
    return {
      photo: null,
      caption: '',
      uploadSuccess: false,
      endText: '',
    };
  },
  methods: {
    onFileChange(event) {
      this.photo = event.target.files[0];
    },
    async uploadPhoto() {
      if (!this.photo) {
        console.log('Photo is required');
        return;
      }

      const formData = new FormData();
      formData.append('photo', this.photo);
      formData.append('caption', this.caption);
      const config = {
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': `Bearer ${token}`,
        },
      };
      try {
        const response = await this.$axios.post(`/photos/`, formData, config);
        console.log('Photo uploaded successfully', response.data);
        this.endText = "Photo uploaded!";
        this.uploadSuccess = true;
      }
      catch (error) {
        const statusCode = error.response.status;
        switch (statusCode) {
          case 401:
            console.error('Access Unauthorized');
            this.endText = "You have to log in to post a photo";
            this.uploadSuccess = true;
            break;
          default:
            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
        }
      }
    },
  },
};
</script>