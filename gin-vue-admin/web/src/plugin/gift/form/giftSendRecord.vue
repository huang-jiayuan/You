
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="赠送者用户ID，关联 user.id:" prop="sendUserId">
          <el-input v-model.number="formData.sendUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="接收者用户ID，关联 user.id:" prop="receiveUserId">
          <el-input v-model.number="formData.receiveUserId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="房间ID(NULL-私聊):" prop="roomId">
          <el-input v-model.number="formData.roomId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="礼物ID:" prop="giftId">
          <el-input v-model.number="formData.giftId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="赠送数量:" prop="sendCount">
          <el-input v-model.number="formData.sendCount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总消耗(虚拟币):" prop="totalPrice">
          <el-input-number v-model="formData.totalPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="总消耗(钻石):" prop="totalDiamond">
          <el-input v-model.number="formData.totalDiamond" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="赠送方式(1-背包,2-直接购买):" prop="sendType">
           <el-select v-model="formData.sendType" placeholder="请选择赠送方式(1-背包,2-直接购买)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in sendtypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="赠送附言:" prop="message">
          <el-input v-model="formData.message" :clearable="true"  placeholder="请输入赠送附言" />
       </el-form-item>
        <el-form-item label="赠送时间:" prop="sendTime">
          <el-date-picker v-model="formData.sendTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="状态(0-失败,1-成功,2-已撤回):" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态(0-失败,1-成功,2-已撤回)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in giftstatueOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="赠送者IP:" prop="clientIp">
          <el-input v-model="formData.clientIp" :clearable="true"  placeholder="请输入赠送者IP" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createGiftSendRecord,
  updateGiftSendRecord,
  findGiftSendRecord
} from '@/plugin/gift/api/giftSendRecord'

defineOptions({
    name: 'GiftSendRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const sendtypeOptions = ref([])
const giftstatueOptions = ref([])
const formData = ref({
            sendUserId: undefined,
            receiveUserId: undefined,
            roomId: undefined,
            giftId: undefined,
            sendCount: undefined,
            totalPrice: 0,
            totalDiamond: undefined,
            sendType: '',
            message: '',
            sendTime: new Date(),
            status: '',
            clientIp: '',
        })
// 验证规则
const rule = reactive({
               sendUserId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               receiveUserId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               giftId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sendCount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalPrice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               totalDiamond : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sendType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGiftSendRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    sendtypeOptions.value = await getDictFunc('sendtype')
    giftstatueOptions.value = await getDictFunc('giftstatue')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createGiftSendRecord(formData.value)
               break
             case 'update':
               res = await updateGiftSendRecord(formData.value)
               break
             default:
               res = await createGiftSendRecord(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
