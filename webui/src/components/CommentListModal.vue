<template>
    <div class="modal fade" tabindex="-1" :id="'listModal' + photoId" aria-labelledby="ModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Comments</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <ul class="list-group">
                        <li v-for="comment in comments" :key="comment.id" class="list-group-item">
                            <div>
                                <strong>{{ comment.publisherName }}</strong>
                            </div>
                            <div>{{ comment.text }}</div>
                            <div v-if="comment.author.userId == this.token">
                                <button @click="deleteComment(comment.photoId, comment.commentId)" class="btn btn-danger btn-sm">Delete</button>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
export default {
    props: {
        photoId: String,
    },
    data() {
        return {
            showModal: false,
            comments: [],
            token: sessionStorage.getItem('authToken'),
        };
    },
    async mounted() {
        this.fetchComments();
    },
    methods: {
        async fetchComments() {
            try {
                const response = await this.$axios.get(`/photos/${this.photoId}/comments/`, {
                    headers: {
                        Authorization: `Bearer ${this.token}`,
                    },
                });
                this.comments = response.data.comments;
            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }

        },
        async deleteComment(postId, commentId) {
            try {
                const response = await this.$axios.delete(`/photos/${postId}/comments/${commentId}`, {
                    headers: {
                        'Authorization': `Bearer ${this.token}`,
                    }
                },);
                location.reload();
            }
            catch (error) {
                console.error(error, "delete-comment: unable to perform action")
            }

        },
    },
};
</script>