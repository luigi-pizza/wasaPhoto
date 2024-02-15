<script>
import PhotoCard from '@/components/PhotoCard.vue';
const token = sessionStorage.getItem('authToken');

export default {

    mounted() {
        if (localStorage.getItem('reloadedstream')) {
            localStorage.removeItem('reloadedstream');
        } else {
            localStorage.setItem('reloadedstream', '1');
            location.reload();
        }
    },
    data() {
        return {
            photoList: [],
            titlePage: "Your stream",
        }
    },
    watch: {
        '$route.params.userId'(newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        },
    },
    async mounted() {
        const userId = this.$route.params.userId;
        this.fetchUserData();
    },
    
    methods: {
        refresh() {
            location.reload();
        },
        async fetchUserData() {
            const userId = this.$route.params.userId;
            try {
                const response = await this.$axios.get(`/stream`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                this.photoList = response.data.posts;
            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 400:
                            console.error('Bad Request:', error.response.data);
                            this.titlePage = "You are not logged in"
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            this.titlePage = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            this.titlePage = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            this.titlePage = "You are not logged in"
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },
    },
    components: {
        PhotoCard,
    },
}
</script>




<template>
    <div class="container mt-5">
        <h1 class="display-4 mb-4">{{ titlePage }}</h1>
        <hr />

        <div class="row row-cols-1 row-cols-md-2 row-cols-lg-3 g-4">
            <PhotoCard v-for="photo in photoList" :key="photo.photoId" :photoId="photo.photoId" :date="photo.creation"
                :authorName="photo.author.username" :likes="photo.numberOfLikes" :caption="photo.caption" :Liked="photo.isliked" class="col mb-4" />
        </div>
    </div>
</template>
  