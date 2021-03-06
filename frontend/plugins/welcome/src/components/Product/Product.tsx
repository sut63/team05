import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Content, Header, Page, pageTheme /*ContentHeader,*/,} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save';
import { Grid, Container } from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import InputAdornment from '@material-ui/core/InputAdornment';
import Swal from 'sweetalert2'; // alert

import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { EntGroupOfAge } from '../../api/models/EntGroupOfAge'; // import interface GroupOfAge
import { EntOfficer} from '../../api/models/EntOfficer'; // import interface Officer

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
    },
    button: {
      margin: theme.spacing(1),
    },
    textField: {
      width: 300,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    formControl: {
      width: 300,
    },
    
  }),
);



export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ระบบประกันสุขภาพ' };
  const api = new DefaultApi();
 
  const [genders, setGenders] = useState<EntGender[]>([]);
  const [groupOfAges, setGroupOfAges] = useState<EntGroupOfAge[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [genderID, setGenderID] = React.useState(Number);
  const [groupOfAgeID, setGroupOfAgeID] = React.useState(Number);
  const [officerID, setOfficerID] = React.useState(Number);


  const [productName, setProductName] = React.useState(String);
  const [productPrice, setProductPrice] = React.useState(Number);
  const [productTime, setProductTime] = React.useState(Number);
  const [productPaymentOfYear, setProductPaymentOfYear] = React.useState(Number);

  useEffect(() => {
    const getGender = async () => {
      const gd = await api.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(gd);
    };
    getGender();

    const getGroupOfAge = async () => {
      const goa = await api.listGroupofage({ limit: 10, offset: 0 });
      setLoading(false);
      setGroupOfAges(goa);
    };
    getGroupOfAge();

    const getOfficer = async () => {
      const ofc = await api.listOfficer({ limit: 10, offset: 0 });
      setLoading(false);
      setOfficers(ofc);
    };
    getOfficer();

    const dataa = localStorage.getItem("officerdata");
    if (dataa) {
    setOfficerID(Number(localStorage.getItem("officerdata")));
    setLoading(false);
    }

    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
      setLoading(false);
      if (jobdata != "พนักงานบริษัทประกันสุขภาพ" ) {
        localStorage.setItem("officerdata",JSON.stringify(null));
        localStorage.setItem("positiondata",JSON.stringify(null));
        history.pushState("","","./Officerlongin");  
        window.location.reload(false);       
      }
      else{
          setOfficerID(Number(localStorage.getItem("officerdata")))
      }
    }
  checkJobPosition();
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


  /*const CreateProduct = async () => {
    if ((productName != null) && (productName != "") && (productPrice != 0) && (productPrice != null) 
    && (productTime != 0) && (productTime != null) && (productPaymentOfYear != null) && (productPaymentOfYear != 0) 
    && (genderID != null) && (groupOfAgeID != null) && (officerID != null)) { 

    const product = {

      gender: genderID,
      groupOfAge: groupOfAgeID,
      officer: officerID,

      productName: String(productName),
      productPrice: Number(productPrice),
      productTime: Number(productTime),
      productPaymentOfYear: Number(productPaymentOfYear),
    };
    console.log(product);

    const res: any = await api.createProduct({ product: product });
            setStatus(true);
            if (res.id != '') {
                setAlert(true);
            }
        }
        else {
            setStatus(true);
            setAlert(false);
        }
    }; */

    const product = {
      gender: genderID,
      groupOfAge: groupOfAgeID,
      officer: officerID,

      productName: String(productName),
      productPrice: Number(productPrice),
      productTime: Number(productTime),
      productPaymentOfYear: Number(productPaymentOfYear),
      };

      function save() {

        const apiUrl = 'http://localhost:8080/api/v1/products';
        const requestOptions = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(product),
        };

        console.log(product); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
          .then(response => response.json())
          .then(data => {
            console.log(data);
            if (data.status == true) {
              //clear();
              Toast.fire({
                icon: 'success',
                title: 'บันทึกข้อมูลสำเร็จ',
              });
            } else {
              checkCaseSaveError(data.error.Name)
            }
          });
      }
  // alert error
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  // check error 
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'product_name':
        alertMessage("error","เป็นตัวใหญ่เท่านั้น (ex: YONGMAN)");
        return;
      case 'product_price':
        alertMessage("error","จำนวนเงินประกันขั้นต่ำ 100,000 บาท");
        return;
      case 'product_time':
        alertMessage("error","ปีไม่ถูกต้อง");
        return;
      case 'product_payment_of_year':
        alertMessage("error","ผ่อนชำระขั้นต่ำ 10,000 บาท");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }


  const GenderhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setGenderID(event.target.value as number);
  };

  const GroupOfAgehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setGroupOfAgeID(event.target.value as number);
  };

  const OfficerhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setOfficerID(event.target.value as number);
  };

  const ProductNamehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductName(event.target.value as string);
  };

  const ProductPricehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductPrice(event.target.value as number);
  };

  const ProductTimehandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductTime(event.target.value as number);
  };

  const ProductPaymentOfYearhandleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setProductPaymentOfYear(event.target.value as number);
  };
  

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`${profile.givenName || 'to Backstage'}`}
        subtitle="ระบบันทึกข้อมูลผลิตภัณฑ์"
      >
       
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              {status ? (
                <div>
                  {alert ? (
                    <Alert severity="success" style={{ width: 300 }}>
                      บันทึกข้อมูลสำเร็จ :)
                    </Alert>
                  ) : (
                    <Alert severity="error" style={{ width: 300 }}>
                       บันทึกข้อมูลไม่สำเร็จ
                    </Alert>
                  )}
                </div>
              ) : null}
            </Grid>
          
            <Grid item xs={4}>
              <div className={classes.paper}>ผลิตภัณฑ์</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  // label="ป้อนชื่อผลิตภัณฑ์"
                  name="productpament"
                  type=" "
                  value={productName}
                  onChange={ProductNamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                />
              </FormControl>
            </Grid>


            <Grid item xs={4}>
              <div className={classes.paper}> เงินประกัน : ฿</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
             <OutlinedInput
                id="outlined-adornment-price"
                value={productPrice}
                onChange={ProductPricehandleChange}
                startAdornment={<InputAdornment position="start">฿</InputAdornment>}
               />
             </FormControl>
            </Grid>


            <Grid item xs={4}>
              <div className={classes.paper}> ระยะเวลา</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
             <OutlinedInput
                id="outlined-adornment-time"
                value={productTime}
                onChange={ProductTimehandleChange}
                startAdornment={<InputAdornment position="start">ปี</InputAdornment>}
               />
             </FormControl>
            </Grid>
            

            <Grid item xs={4}>
              <div className={classes.paper}> ผ่อนชำระ / ปี</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
             <OutlinedInput
                id="outlined-adornment-payment"
                value={productPaymentOfYear}
                onChange={ProductPaymentOfYearhandleChange}
                startAdornment={<InputAdornment position="start">฿</InputAdornment>}
               />
             </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="gender"
                  value={genderID}
                  onChange={GenderhandleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.genderName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>กลุ่มอายุ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="groupofagr"
                  value={groupOfAgeID}
                  onChange={GroupOfAgehandleChange}
                >
                  {groupOfAges.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.groupOfAgeAge}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ผู้บันทึกข้อมูล</div>
            </Grid>
            <Grid item xs={8}>
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


            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="small"
                className={classes.button}
                startIcon={<SaveIcon />}
                onClick={() => {
                  save();
                }}
              >
                Save
              </Button>
             
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
}