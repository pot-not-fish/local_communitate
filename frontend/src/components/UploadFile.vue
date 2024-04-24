<template>
    <el-input
        v-model="input"
        style="width: 240px"
        placeholder="请输入目的地址"
        clearable
    >
        <template #prepend>Host</template>
    </el-input>
    <el-button type="primary">确定传输</el-button>

    <el-upload
        class="upload-demo"
        drag
        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
        multiple
    >
        <el-icon class="el-icon--upload">
            <upload-filled />
        </el-icon>
        <div class="el-upload__text">
            将文件拖拽到此处 <em>点击上传文件</em>
        </div>
        <template #tip>
            <div class="el-upload__tip">
                jpg/png files with a size less than 500kb
            </div>
        </template>
    </el-upload>
</template>
  
<script setup lang="ts">
    import { UploadFilled } from '@element-plus/icons-vue'

    import { ref } from 'vue'
    const input = ref('')
</script>
