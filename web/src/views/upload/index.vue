<template>
  <div class="page-content-container">
    <el-form  label-width="100px">
      <el-row>
        <el-col :span="24">
          <h3>选择上传文件：</h3>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="24" type="flex">
            <el-upload  :action="UploadUrl()"  :on-success="UploadSuccess">
              <el-button size="small" type="primary" >点击上传</el-button>
            </el-upload>
        </el-col>
        </el-row>

      <el-row>
        <el-col :span="24">
        
       <el-table :key='tableKey' :data="list" border fit highlight-current-row
        width= "100%" max-height= "580" >
        <el-table-column align="center" :label="'行号'" width="80"   prop="RowIndex">
            <template slot-scope="scope">
            <span>{{scope.row.RowIndex}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'磅单编号'" width="180"   prop="NumberFlag">
            <template slot-scope="scope">
            <span>{{scope.row.NumberFlag}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'首次称重日期'" width="140"   prop="Date">
            <template slot-scope="scope">
            <span>{{scope.row.Date}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'毛重时间'" width="140"   prop="Time">
            <template slot-scope="scope">
            <span>{{scope.row.Time}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'车号'" width="100"   prop="CarNumber">
            <template slot-scope="scope">
            <span>{{scope.row.CarNumber}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'货物名称'" width="100"   prop="ShopNumber">
            <template slot-scope="scope">
            <span>{{scope.row.ShopNumber}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'毛重'" width="80"   prop="GrossWeight">
            <template slot-scope="scope">
            <span>{{scope.row.GrossWeight}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'皮重'" width="80"   prop="TareWeight">
            <template slot-scope="scope">
            <span>{{scope.row.TareWeight}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'净重'" width="80"   prop="NetWeight">
            <template slot-scope="scope">
            <span>{{scope.row.NetWeight}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'单价'" width="80"   prop="UnitPrice">
            <template slot-scope="scope">
            <span>{{scope.row.UnitPrice}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" :label="'收货单位'" prop="TakeDept">
            <template slot-scope="scope">
            <span>{{scope.row.TakeDept}}</span>
            </template>
        </el-table-column>
        
     </el-table>     
     </el-col>
    </el-row>

    </el-form>
  </div>
</template>

<script>
import Vue from 'vue'
import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default {
  name: 'test',
  data() {
    return {
      tableKey: 1,
      list: null
    }
  },
  components: {
  },
  created() {
  },
  methods: {
    UploadUrl() {
      return 'http://localhost:8080/Upload/UpFile'
    },
    UploadSuccess(rsp) {
      this.list = rsp
    }
  }
}
</script>

<style>
  .uploader-example {
    width: 880px;
    padding: 15px;
    margin: 40px auto 0;
    font-size: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, .4);
  }
  .uploader-example .uploader-btn {
    margin-right: 4px;
  }
  .uploader-example .uploader-list {
    max-height: 440px;
    overflow: auto;
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
