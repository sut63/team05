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
import { EntCategory } from '../../api/models/EntCategory';
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
  const [categorys, setCategorys] = useState<EntCategory[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [productid, setProductid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [categoryid, setCategoryid] = useState(Number);
  const [officerid, setOfficerid] = useState(Number);
  const [inquiry_messages, setInquiryMessages] = useState(String);
  const [inquiry_phone_messages, setInquiryPhoneMessages] = useState(String);
  //const [inquiryTimeMessages, setInquiryTimeMessages] = useState(String);
  
  useEffect(() => {
 
    const getProducts = async () => {
 
      const p = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(p);
    };
    getProducts();
 
    const getMembers = async () => {
 
        const u = await api.listMember({});
        setLoading(false);
        setMembers(u);
    };
    getMembers();
 
    const getCategorys = async () => {
 
     const cg = await api.listCategory({ limit: 10, offset: 0 });
       setLoading(false);
       setCategorys(cg);
     };
     getCategorys();

     const getOfficers = async () => {
 
      const of = await api.listOfficer({ limit: 10, offset: 0 });
        setLoading(false);
        setOfficers(of);
      };
      getOfficers();


      const data = localStorage.getItem("memberdata");
      if (data) {
      setMemberid(Number(localStorage.getItem("memberdata")));
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

  const CreateInquiry = async ()=>{
    if ((productid != null) && (categoryid != null) && (officerid != null) && (inquiry_messages != null) && (inquiry_messages != "") && (inquiry_phone_messages != null) && (inquiry_phone_messages != "")) {
    const inquiry ={


         inquiryMessages      : inquiry_messages,
         inquiryPhoneMessages      : inquiry_phone_messages,
         //inquiryTimeMessages     : inquiryTimeMessages + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         productID          : productid,
         memberID        : memberid,
         categoryID    : categoryid,
         officerID         : officerid,
      };
      console.log(inquiry);
      const res: any = await api.createInquiry({ inquiry: inquiry });
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
      /*const timer = setTimeout(() => {
        setStatus(false);
        }, 10000);*/
    };
 
   const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
   setProductid(event.target.value as number);
    };

   const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setMemberid(event.target.value as number);
    };

    const category_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setCategoryid(event.target.value as number);
     };

     const officer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setOfficerid(event.target.value as number);
     };

     const inquiryMessages_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInquiryMessages(event.target.value as string);
     };

     const inquiryPhoneMessages_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInquiryPhoneMessages(event.target.value as string);
     };

     /*const inquiryTimeMessages_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInquiryTimeMessages(event.target.value as string);
     };*/

     return (
        <Page theme={pageTheme.app}>
          <Header style={HeaderCustom} title={`ระบบติดต่อสอบถามข้อมูล`}>
          </Header>
          <Content>
          <ContentHeader title="บันทึกการสอบถามข้อมูล">
  
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกสำเร็จ!
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
                  
                <Grid item xs={4}>
                <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={8}>
                  <FormControl
                    className={classes.formControl }
                    variant="outlined"
                  >
                  <TextField
                    id="user"
                    //label="ชื่อสมาชิก"
                    type="string"
                    size="medium"
                    variant="outlined"
                    value={members.filter((filter: EntMember) => filter.id == memberid).map((item: EntMember) => `${item.memberName}`)}
                    style={{ width: 300  }}
                  />
                  </FormControl>
              </Grid>

              <Grid item xs={4}>
              <div className={classes.paper}>รายการสอบถาม</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>เลือกรายการสอบถาม</InputLabel>
                  <Select
                    name="category"
                    value={categoryid || ''} // (undefined || '') = ''
                    onChange={category_id_handleChange}
                  >
                    {categorys.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.categoryName}
                        </MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
            </Grid>
  
              <Grid item xs={4}>
                <div className={classes.paper}>ชื่อผลิตภัณฑ์</div>
              </Grid>
              <Grid item xs={8}>
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
  
              <Grid item xs={4}>
                <div className={classes.paper}>พนักงานบริษัทประกันสุขภาพ</div>
              </Grid>
              <Grid item xs={8}>
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

              <Grid item xs={4}>
              <div className={classes.paper}>เบอร์โทรติดต่อกลับ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="inquiry_phone_messages"
                  multiline
                  rows={1}
                  value={inquiry_phone_messages}
                  type="string"
                  onChange={inquiryPhoneMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>

              <Grid item xs={4}>
              <div className={classes.paper}>รายละเอียด</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  id="outlined-number"
                  name="inquiry_messages"
                  multiline
                  rows={4}
                  value={inquiry_messages}
                  type="string"
                  onChange={inquiryMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>
    
                <Grid item xs={4}></Grid>
                <Grid item xs={8}>
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  startIcon={<SaveIcon />}
                  onClick={() => {
                    CreateInquiry();
                  }}
                >
                  บันทึกการสอบถาม
                </Button>
                </Grid>
              </Grid>
            </Container>
          </Content>
        </Page>
      );
  };