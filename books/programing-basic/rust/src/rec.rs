pub fn add_to_each<T>(x: T, list: Vec<Vec<T>>) -> Vec<Vec<T>>
where
    T: Clone,
{
    match list {
        l if l.len() == 0 => vec![],
        l => {
            let mut first = l[0].clone();
            first.insert(0, x.clone());
            let rest = l[1..].to_vec();
            vec![first]
                .into_iter()
                .chain(add_to_each(x, rest))
                .collect()
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_add_to_each() {
        let list = vec![vec![1], vec![2], vec![3]];
        let expected = vec![vec![0, 1], vec![0, 2], vec![0, 3]];
        assert_eq!(add_to_each(0, list), expected);
    }
    //#[test]
    //fn test_prefix() {
    //let list = vec![1, 2, 3, 4, 5];
    //let expected = vec![
    //vec![1],
    //vec![1, 2],
    //vec![1, 2, 3],
    //vec![1, 2, 3, 4],
    //vec![1, 2, 3, 4, 5],
    //];
    //assert_eq!(prefix(&list), expected);
    //}
}
