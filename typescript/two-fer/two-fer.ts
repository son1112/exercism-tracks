class TwoFer {
    static twoFer( name?: string ) {
        var subject = name ? name : "you"
        var message = `One for ${subject}, one for me.`

        return message
    }
}

export default TwoFer
