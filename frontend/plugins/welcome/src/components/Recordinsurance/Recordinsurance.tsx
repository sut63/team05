import React, { FC, useEffect,useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core'
import { DefaultApi } from '../../api/apis';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { EntProduct } from '../../api/models/EntProduct';
import { EntMember } from '../../api/models/EntMember';
import { EntHospital } from '../../api/models/EntHospital';
import { EntOfficer } from '../../api/models/EntOfficer';
import { EntAmountpaid } from '../../api/models/EntAmountpaid';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import InputAdornment from '@material-ui/core/InputAdornment';
import { Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';
import { ContentHeader } from '@backstage/core';
import { Link } from '@material-ui/core';
import { Theme, createStyles } from '@material-ui/core/styles';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

export default function Create() {
    const classes = useStyles();
    const api = new DefaultApi();
   
    const [members, setMembers] = useState<EntMember[]>([]);
    const [products, setProducts] = useState<EntProduct[]>([]);
    const [amountpaids, setAmountpaids] = useState<EntAmountpaid[]>([]);
    const [officers, setOfficers] = useState<EntOfficer[]>([]);
    const [hospitals, setHospitals] = useState<EntHospital[]>([]);
  
    const [status, setStatus] = useState(false);
    const [alert, setAlert] = useState(true);
    const [loading, setLoading] = useState(true);
  
    const [memberid, setMemberid] = React.useState(Number);
    const [productid, setProductid] = React.useState(Number);
    const [amountpaidid, setAmountpaidid] = React.useState(Number);
    const [officerid, setOfficerid] = React.useState(Number);
    const [hospitalid, setHospitalid] = useState(Number);
 
  const [recordinsurance_time, setRecordinsuranceTime] = useState(String);
  
 
  useEffect(() => {
 
    const getProducts = async () => {
 
      const cn = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(cn);
    };
    getProducts();
 
    const getMembers = async () => {
 
    const u = await api.listMember({ limit: 10, offset: 0 });
      setLoading(false);
      setMembers(u);
    };
    getMembers();
 
    const getHospitals = async () => {
 
     const s = await api.listHospital({ limit: 10, offset: 0 });
       setLoading(false);
       setHospitals(s);
     };
     getHospitals();

     const getOfficers = async () => {
 
      const st = await api.listOfficer({ limit: 10, offset: 0 });
        setLoading(false);
        setOfficers(st);
      };
      getOfficers();

      const getAmountpaids = async () => {
 
        const am = await api.listAmountpaid({ limit: 10, offset: 0 });
          setLoading(false);
          setAmountpaids(am);
        };
        getAmountpaids();
 
  }, [loading]);
 
  /*const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });*/

  const CreateRecordinsurance = async () => {
    if ((recordinsurance_time != null) && (recordinsurance_time != "") && (amountpaidid != null) && (hospitalid != null) && (memberid != null) && (officerid != null) && (productid != null)){
      const recordinsurance = {

         amountpaidID      : amountpaidid,
         hospitalID        : hospitalid,
         memberID          : memberid,
         officerID         : officerid,
         productID         : productid,
         recordinsuranceTime    : recordinsurance_time + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
      }
      console.log(recordinsurance);
    
    const res:any = await api.createRecordinsurance({ recordinsurance : recordinsurance});
            setStatus(true);
            if (res.id != '') {
                setAlert(true);
            }
        }
        else {
            setStatus(true);
            setAlert(false);
        }
    };
 
   const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
   setProductid(event.target.value as number);
    };

   const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMemberid(event.target.value as number);
    };

    const hospital_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setHospitalid(event.target.value as number);
     };

     const officer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setOfficerid(event.target.value as number);
     };

     const amountpaid_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAmountpaidid(event.target.value as number);
       };

       const recordinsuranceTime_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setRecordinsuranceTime(event.target.value as string);
      };

    return (
      <Page theme={pageTheme.tool}>
        <Header style={HeaderCustom} title={`ระบบบันทึกข้อมูลสิทธิประกันสุขภาพ`}>
        </Header>
        <Content>
        <ContentHeader title="ทำการบันทึกข้อมูลสิทธิประกันสุขภาพ" >

            {status ? (
                <div>
                    {alert ? (
                        <Alert severity="success">
                        บันทึกเรียบร้อย!
                        </Alert>
                    ) : (
                        <Alert severity="warning" style={{ marginTop: 20 }}>
                        บันทึกไม่สำเร็จ!
                        </Alert>
                    )}
                </div>
            ) : null}
        </ContentHeader>

          <Container maxWidth="sm">
          <Grid container spacing={3}>
          <Grid item xs={10}></Grid>
          <Grid item xs={3}>
                <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>สมาชิกระบบประกันสุขภาพ</InputLabel>
                  <Select
                    name="member"
                    value={memberid || ''} // (undefined || '') = ''
                    onChange={member_id_handleChange}
                  >
                    {members.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.memberName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
                <div className={classes.paper}>ผลิตภัณฑ์</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกผลิตภัณฑ์</InputLabel>
                  <Select
                    name="product"
                    value={productid || ''} // (undefined || '') = ''
                    onChange={product_id_handleChange}
                  >
                    {products.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.productName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
              <div className={classes.paper}> เงินที่ประกันภัยจ่าย </div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                    name="amountpaid"
                    value={amountpaidid || ''} // (undefined || '') = ''
                    onChange={amountpaid_id_handleChange}
                  >
                    {amountpaids.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.amountpaidMoney}
                        </MenuItem>
                      );
                    })}
                  </Select>
              </FormControl>
            </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>โรงพยาบาล</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกโรงพยาบาล</InputLabel>
                  <Select
                    name="hospital"
                    value={hospitalid || ''} // (undefined || '') = ''
                    onChange={hospital_id_handleChange}
                  >
                    {hospitals.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.hospitalName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>พนักงานบริษัทประกันสุขภาพที่แนะนำ</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกพนักงานบริษัทประกันสุขภาพ</InputLabel>
                  <Select
                    name="officer"
                    value={officerid || ''} // (undefined || '') = ''
                    onChange={officer_id_handleChange}
                  >
                    {officers.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.officerName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
               <div className={classes.paper}>เวลาที่ทำการบันทึกข้อมูล</div>
              </Grid>
              <Grid item xs={9}>
               <form className={classes.container} noValidate>
                <TextField
                  label="เวลาที่บันทึก"
                  name="recordinsurance_time"
                  type="datetime-local"
                  value={recordinsurance_time || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                   onChange={recordinsuranceTime_handleChange}
                 />
               </form>
             </Grid>

              <Grid item xs={3}></Grid>
              <Grid item xs={7}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateRecordinsurance();
                }}
              >
                บันทึกข้อมูลสิทธิประกันสุขภาพ
              </Button>
            </Grid>
          </Grid>
          </Container>
        </Content>
      </Page>
    );
  };