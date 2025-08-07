
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="礼物名称" prop="giftName">
  <el-input v-model="searchInfo.giftName" placeholder="搜索条件" />
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="gift_GiftInfo" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="gift_GiftInfo" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="gift_GiftInfo" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="giftId"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="礼物唯一ID" prop="giftId" width="120" />

            <el-table-column align="left" label="礼物名称" prop="giftName" width="120" />

            <el-table-column align="left" label="礼物类型(1-普通,2-特效,3-稀有,4-自定义)" prop="giftType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.giftType,gifttypeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="虚拟币价格" prop="price" width="120" />

            <el-table-column align="left" label="钻石价格(可为空)" prop="diamondPrice" width="120" />

            <el-table-column align="left" label="礼物图标URL" prop="iconUrl" width="120" />

            <el-table-column align="left" label="动画URL(特效礼物)" prop="animationUrl" width="120" />

            <el-table-column align="left" label="排序权重" prop="weight" width="120" />

            <el-table-column align="left" label="是否热门(1-是,0-否)" prop="isHot" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.isHot,ishotOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="是否限时(1-是,0-否)" prop="isLimit" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.isLimit,islimitOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="状态(0-下架,1-上架,2-待审核)" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,giftstatusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="修改时间" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateGiftInfoFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="礼物唯一ID:" prop="giftId">
    <el-input v-model.number="formData.giftId" :clearable="true" placeholder="请输入礼物唯一ID" />
</el-form-item>
             <el-form-item label="礼物名称:" prop="giftName">
    <el-input v-model="formData.giftName" :clearable="true" placeholder="请输入礼物名称" />
</el-form-item>
             <el-form-item label="礼物类型(1-普通,2-特效,3-稀有,4-自定义):" prop="giftType">
    <el-select v-model="formData.giftType" placeholder="请选择礼物类型(1-普通,2-特效,3-稀有,4-自定义)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in gifttypeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="虚拟币价格:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
             <el-form-item label="钻石价格(可为空):" prop="diamondPrice">
    <el-input v-model.number="formData.diamondPrice" :clearable="true" placeholder="请输入钻石价格(可为空)" />
</el-form-item>
             <el-form-item label="礼物图标URL:" prop="iconUrl">
    <el-input v-model="formData.iconUrl" :clearable="true" placeholder="请输入礼物图标URL" />
</el-form-item>
             <el-form-item label="动画URL(特效礼物):" prop="animationUrl">
    <el-input v-model="formData.animationUrl" :clearable="true" placeholder="请输入动画URL(特效礼物)" />
</el-form-item>
             <el-form-item label="排序权重:" prop="weight">
    <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入排序权重" />
</el-form-item>
             <el-form-item label="是否热门(1-是,0-否):" prop="isHot">
    <el-select v-model="formData.isHot" placeholder="请选择是否热门(1-是,0-否)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in ishotOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="是否限时(1-是,0-否):" prop="isLimit">
    <el-select v-model="formData.isLimit" placeholder="请选择是否限时(1-是,0-否)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in islimitOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="状态(0-下架,1-上架,2-待审核):" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态(0-下架,1-上架,2-待审核)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in giftstatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="修改时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="礼物唯一ID">
    {{ detailFrom.giftId }}
</el-descriptions-item>
                 <el-descriptions-item label="礼物名称">
    {{ detailFrom.giftName }}
</el-descriptions-item>
                 <el-descriptions-item label="礼物类型(1-普通,2-特效,3-稀有,4-自定义)">
    {{ detailFrom.giftType }}
</el-descriptions-item>
                 <el-descriptions-item label="虚拟币价格">
    {{ detailFrom.price }}
</el-descriptions-item>
                 <el-descriptions-item label="钻石价格(可为空)">
    {{ detailFrom.diamondPrice }}
</el-descriptions-item>
                 <el-descriptions-item label="礼物图标URL">
    {{ detailFrom.iconUrl }}
</el-descriptions-item>
                 <el-descriptions-item label="动画URL(特效礼物)">
    {{ detailFrom.animationUrl }}
</el-descriptions-item>
                 <el-descriptions-item label="排序权重">
    {{ detailFrom.weight }}
</el-descriptions-item>
                 <el-descriptions-item label="是否热门(1-是,0-否)">
    {{ detailFrom.isHot }}
</el-descriptions-item>
                 <el-descriptions-item label="是否限时(1-是,0-否)">
    {{ detailFrom.isLimit }}
</el-descriptions-item>
                 <el-descriptions-item label="状态(0-下架,1-上架,2-待审核)">
    {{ detailFrom.status }}
</el-descriptions-item>
                 <el-descriptions-item label="创建时间">
    {{ detailFrom.createdAt }}
</el-descriptions-item>
                 <el-descriptions-item label="修改时间">
    {{ detailFrom.updatedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createGiftInfo,
  deleteGiftInfo,
  deleteGiftInfoByIds,
  updateGiftInfo,
  findGiftInfo,
  getGiftInfoList
} from '@/plugin/gift/api/giftInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'GiftInfo'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const gifttypeOptions = ref([])
const ishotOptions = ref([])
const giftstatusOptions = ref([])
const islimitOptions = ref([])
const formData = ref({
            giftId: undefined,
            giftName: '',
            giftType: '',
            price: 0,
            diamondPrice: undefined,
            iconUrl: '',
            animationUrl: '',
            weight: undefined,
            isHot: '',
            isLimit: '',
            status: '',
            createdAt: new Date(),
            updatedAt: new Date(),
        })



// 验证规则
const rule = reactive({
               giftName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               giftType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getGiftInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    gifttypeOptions.value = await getDictFunc('gifttype')
    ishotOptions.value = await getDictFunc('ishot')
    giftstatusOptions.value = await getDictFunc('giftstatus')
    islimitOptions.value = await getDictFunc('islimit')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteGiftInfoFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const giftIds = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          giftIds.push(item.giftId)
        })
      const res = await deleteGiftInfoByIds({ giftIds })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === giftIds.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateGiftInfoFunc = async(row) => {
    const res = await findGiftInfo({ giftId: row.giftId })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteGiftInfoFunc = async (row) => {
    const res = await deleteGiftInfo({ giftId: row.giftId })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        giftId: undefined,
        giftName: '',
        giftType: '',
        price: 0,
        diamondPrice: undefined,
        iconUrl: '',
        animationUrl: '',
        weight: undefined,
        isHot: '',
        isLimit: '',
        status: '',
        createdAt: new Date(),
        updatedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createGiftInfo(formData.value)
                  break
                case 'update':
                  res = await updateGiftInfo(formData.value)
                  break
                default:
                  res = await createGiftInfo(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findGiftInfo({ giftId: row.giftId })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
