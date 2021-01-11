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

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [hospitals, setHospitals] = useState<EntHospital[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [productid, setProductid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [hospitalid, setHospitalid] = useState(Number);
  const [officerid, setOfficerid] = useState(Number);
  const [insurance_addresss, setInsuranceAddress] = useState(String);
  const [insurance_insurers, setInsuranceInsurer] = useState(String);
  const [insuranceTimeBuys, setInsuranceTimeBuy] = useState(String);
  const [insuranceTimeFirstpays, setInsuranceTimeFirstpay] = useState(String);
  const [productPrice, setProductPrice] = useState(String);
  const [productPaymentOfYear, setProductPaymentOfYear] = useState(String);
  
 
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
 
  }, [loading]);
 
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const CreateInsurance = async () => {
      const insurance = {

         insuranceAddress      : insurance_addresss,
         insuranceInsurer     : insurance_insurers,
         insuranceTimeBuy     : insuranceTimeBuys + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         insuranceTimeFirstpay  : insuranceTimeFirstpays  + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         productID          : productid,
         memberID        : memberid,
         hospitalID    : hospitalid,
         officerID         : officerid,
      }
      console.log(insurance);

    const res:any = await api.createInsurance({ insurance : insurance});
    setStatus(true);
    if (res !== undefined) {
      Toast.fire({
        icon: 'success',
        title: 'บันทึกข้อมูลสำเร็จ',
      });
    } else {
      Toast.fire({
        icon: 'error',
        title: 'บันทึกข้อมูลไม่สำเร็จ',
      });
    }
 
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
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

     const insuranceAddress_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setInsuranceAddress(event.target.value as string);
     };

     const insuranceInsurer_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setInsuranceInsurer(event.target.value as string);
     };

     const insuranceTimeBuy_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setInsuranceTimeBuy(event.target.value as string);
    };

    const insuranceTimeFirstpay_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setInsuranceTimeFirstpay(event.target.value as string);
    };
    
    const productPrice_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProductPrice(event.target.value as string);
    };
    
    const productPaymentOfYear_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProductPaymentOfYear(event.target.value as string);
    };

    return (
      <Page theme={pageTheme.tool}>
        <Header style={HeaderCustom} title={`ระบบการซื้อประกันสุขภาพ`}>
          <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
          <div style={{ marginLeft: 10 }}>Teerapat Saiprom</div>
        </Header>
        <Content>
          <Container maxWidth="sm">
            <Grid container spacing={3}>
              <Grid item xs={12}></Grid>
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
              <div className={classes.paper}> ราคา </div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel> </InputLabel>
                <Select
                    name="product"
                    value={productid || ''} // (undefined || '') = ''
                    onChange={product_id_handleChange}
                    startAdornment={<InputAdornment position="start">฿</InputAdornment>}
                  >
                  {products.map((item: EntProduct) => (
                    <MenuItem value={item.id}>{item.productPrice}</MenuItem>
                  ))
                  }
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
              <div className={classes.paper}>ที่อยู่ของผู้ซื้อ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="insurance_addresss"
                  multiline
                  rows={4}
                  value={insurance_addresss}
                  type="string"
                  onChange={insuranceAddress_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ข้อมูลของผู้รับผลประโยชน์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="insurance_insurer"
                  multiline
                  rows={4}
                  value={insurance_insurers}
                  type="string"
                  onChange={insuranceInsurer_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>

              <Grid item xs={3}>
                <div className={classes.paper}>เวลาที่ซื้อประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={9}>
                <form className={classes.container} noValidate>
                  <TextField
                    label="เลือกเวลาที่ซื้อ"
                    name="insurance_timebuy"
                    type="datetime-local"
                    value={insuranceTimeBuys || ''} // (undefined || '') = ''
                    className={classes.textField}
                    InputLabelProps={{
                      shrink: true,
                    }}
                    onChange={insuranceTimeBuy_handleChange}
                  />
                </form>
              </Grid>
  
              <Grid item xs={3}>
                <div className={classes.paper}>วันที่ต้องการจ่ายงวดแรก</div>
              </Grid>
              <Grid item xs={9}>
                <form className={classes.container} noValidate>
                  <TextField
                    label="เลือกเวลา"
                    name="insurance_timefirstpay"
                    type="datetime-local"
                    value={insuranceTimeFirstpays || ''} // (undefined || '') = ''
                    className={classes.textField}
                    InputLabelProps={{
                      shrink: true,
                    }}
                    onChange={insuranceTimeFirstpay_handleChange}
                  />
                </form>
              </Grid>
  
              <Grid item xs={4}></Grid>
              <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateInsurance();
                }}
              >
                บันทึกการซื้อประกันสุขภาพ
              </Button>
            </Grid>
          </Grid>
          </Container>
        </Content>
      </Page>
    );
  };
  