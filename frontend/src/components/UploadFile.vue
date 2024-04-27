<template>
    <el-input
        v-model="input"
        style="width: 240px"
        placeholder="请输入目的地址"
        clearable
    >
        <template #prepend>Host</template>
    </el-input>
    <el-button type="primary" @click="send">确定传输</el-button>

    <el-upload
        class="upload-demo"
        ref="uploadRef"
        action
        drag
        :limit="1"
        :http-request="fileUpload"
        :auto-upload=false
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
    import { Send } from '../../wailsjs/go/main/App'
    import type { UploadInstance } from 'element-plus'

    import { ref } from 'vue'
    const input = ref('')
    const uploadRef = ref<UploadInstance>()

    function send() {
        uploadRef.value!.submit() // uploadfile
    }

    function fileUpload(file: any) {
        console.log(file.file.name)
        let reader = new FileReader()
        reader.readAsArrayBuffer(file.file)
        reader.onload = (e) => {
            if (e.target !== null) {
                let data = e.target.result as ArrayBuffer
                let uint8array = new Uint8Array(data)
                let arr = Array.from(uint8array)
                Send(input.value, file.file.name, arr).then(result => {
                    alert(result)
                })
            }
        }
    }
</script>
