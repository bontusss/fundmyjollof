const useIdGen = computed(() => {
   return Math.random().toString(36).slice(2)
})

const useDateFormat = (inputDate) => {
   const [day, month, year] = inputDate.split('-');
   const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
   const suffixes = ['st', 'nd', 'rd', 'th'];
   
   const dayNum = parseInt(day);
   const daySuffix = dayNum > 3 ? suffixes[3] : suffixes[dayNum - 1];
   
   const formattedDate = `${dayNum}${daySuffix} ${months[parseInt(month) - 1]}, ${year}`;
   
   return formattedDate;
}

export {
  
}