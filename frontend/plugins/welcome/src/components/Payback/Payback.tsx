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
import { EntBank } from '../../api/models/EntBank';
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
  const [banks, setBanks] = useState<EntBank[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [productid, setProductid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [bankid, setBankid] = useState(Number);
  const [officerid, setOfficerid] = useState(Number);
  const [officerID, setOfficerID] = React.useState(Number);
  const [payback_accountnumber, setPaybackAccountnumber] = useState(String);
  const [payback_transfertime, setPaybackTransfertime] = useState(String);
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
 
    const getBanks = async () => {
 
     const s = await api.listBank({ limit: 10, offset: 0 });
       setLoading(false);
       setBanks(s);
     };
     getBanks();

     const getOfficers = async () => {
 
      const st = await api.listOfficer({ limit: 10, offset: 0 });
        setLoading(false);
        setOfficers(st);
      };
      getOfficers();

      const dataa = localStorage.getItem("officerdata");
    if (dataa) {
    setOfficerID(Number(localStorage.getItem("officerdata")));
    setLoading(false);
    }
 
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

  const CreatePayback = async () => {
    if ((bankid != null) && (memberid != null) && (officerid != null) && (productid != null) && (payback_accountnumber != "") && (payback_transfertime != "")) {
      const payback = {

         bankID    : bankid, 
         memberID        : memberid,
         officerID         : officerID,
         paybackAccountnumber      : payback_accountnumber,
         paybackTransfertime  : payback_transfertime  + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         productID          : productid,     

         
      }
      console.log(payback);

    const res:any = await api.createPayback({ payback : payback});
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      //window.location.reload(false);
    } 
  }
  else {
      setStatus(true)
      setAlert(false);
    }
 
    
  };
 
   const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
   setProductid(event.target.value as number);
    };

   const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMemberid(event.target.value as number);
    };

    const bank_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setBankid(event.target.value as number);
     };

     const officer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setOfficerid(event.target.value as number);
     };

     const paybackAccountnumber_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setPaybackAccountnumber(event.target.value as string);
     };

    const paybackTransfertime_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setPaybackTransfertime(event.target.value as string);
    };
    
    const productPrice_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProductPrice(event.target.value as string);
    };
    
    const productPaymentOfYear_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setProductPaymentOfYear(event.target.value as string);
    };

    return (
      <Page theme={pageTheme.theme}>
        <Header style={HeaderCustom} title={`ระบบย่อย ระบบคืนทุนประกัน`}>

        </Header>
        <Content>
          
        <ContentHeader title="ทำการบันทึกข้อมูลคืนทุนประกัน" >

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
              <Grid item xs={12}></Grid>


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
                <div className={classes.paper}>ธนาคาร</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกธนาคาร</InputLabel>
                  <Select
                    name="bank"
                    value={bankid || ''} // (undefined || '') = ''
                    onChange={bank_id_handleChange}
                  >
                    {banks.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.bankType}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>             

              <Grid item xs={3}>
              <div className={classes.paper}>เลขบัญชี</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="payback_accountnumber"
                  multiline
                  rows={4}
                  value={payback_accountnumber}
                  type="string"
                  onChange={paybackAccountnumber_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
                <div className={classes.paper}>พนักงานที่ทำรายการ</div>
              </Grid>
              <Grid item xs={9}>
              <TextField id="outlined-basic" 
              style={{ width: 300}}
              name = "officer"
              variant="outlined"
              value={officers.filter((filter: EntOfficer) => filter.id == officerID).map((item: EntOfficer) => `${item.officerName}`)}
              className={classes.textField}
              InputLabelProps={{
                shrink: true,
              }}
              
              />
              </Grid>

              
              <Grid item xs={3}>
                <div className={classes.paper}>วันที่ทำรายการ</div>
              </Grid>
              <Grid item xs={9}>
                <form className={classes.container} noValidate>
                  <TextField
                    label="เลือกเวลา"
                    name="payback_transfertime"
                    type="datetime-local"
                    value={payback_transfertime || ''} // (undefined || '') = ''
                    className={classes.textField}
                    InputLabelProps={{
                      shrink: true,
                    }}
                    onChange={paybackTransfertime_handleChange}
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
                  CreatePayback();
                }}
              >
                บันทึกการคืนทุนประกัน
              </Button>
            </Grid>
          </Grid>
          </Container>
        </Content>
      </Page>
    );
  };