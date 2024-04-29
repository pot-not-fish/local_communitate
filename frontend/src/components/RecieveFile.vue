<template>
    <el-table :data="tableData" stripe style="width: 100%">
        <el-table-column prop="date" label="接收时间" width="180" />
        <el-table-column prop="name" label="文件名" width="180" />
        <el-table-column prop="state" label="状态" />
    </el-table>
</template>
  
<script lang="ts" setup>
    import { EventsOn } from '../../wailsjs/runtime/runtime'
    import { ref } from 'vue'

    interface TableRow {  
        date: string;  
        name: string;  
        state: string; // 假设 state 是字符串类型，根据实际情况调整
    }  

    const tableData = ref<TableRow[]>([]);  

    let eventName = "upload_list"
    let callback = (date: any, name: any, state: any) => {
        tableData.value.push({  
            date: date,  
            name: name,  
            state: state,
        })
    }

    const register = EventsOn(eventName, callback)
</script>